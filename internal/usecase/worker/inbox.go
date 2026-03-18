package worker

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"glintfed.org/ent"
	"glintfed.org/ent/profile"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
)

type ProfileRemover interface {
	RemoteProfile(ctx context.Context, profile *ent.Profile) error
}

type InboxUsecase struct {
	client *data.Client

	pr ProfileRemover
	ah *ActivityHandler
}

type InboxPayload struct {
	Raw   string
	ID    string  `json:"id"`
	Type  *string `json:"type,omitempty"`
	Actor *string `json:"actor,omitempty"`

	Object *InboxPayloadObject `json:"object,omitempty"`
}

type InboxPayloadObject struct {
	ID           string  `json:"id"`
	Type         string  `json:"type"`
	AttributedTo *string `json:"attributedTo"`
}

func NewInboxUsecase(client *data.Client) *InboxUsecase {
	return &InboxUsecase{
		client: client,
		pr:     NewDeletePipeline(client),
		ah:     NewActivityHandler(client),
	}
}

func (inbox *InboxUsecase) Delete(ctx context.Context, header http.Header, payload InboxPayload) error {
	if header.Get("signature") == "" || header.Get("date") == "" {
		return errors.New("missing required field in header")
	}

	if err := inbox.verifySignature(ctx, header, payload, "/f/inbox"); err != nil {
		if !ent.IsNotFound(err) {
			return fmt.Errorf("failed to verify signature: %w", err)
		} else {
			slog.WarnContext(ctx, "ignored missing profile", logs.ErrAttr(err))
		}
		return nil
	}

	if payload.Object.Type == "Person" && payload.Actor != nil && *payload.Actor == payload.Object.ID {
		profile, err := inbox.client.Ent.Profile.Query().
			Where(profile.DomainNotNil(), profile.StatusIsNil(), profile.RemoteURL(payload.Object.ID)).
			First(ctx)
		if err != nil {
			return fmt.Errorf("failed to get profile: %w", err)
		}

		return inbox.pr.RemoteProfile(ctx, profile)
	} else {
		return inbox.ah.Delete(ctx, header, payload)
	}
}

func (inbox *InboxUsecase) Inbox(ctx context.Context, header http.Header, payload InboxPayload) error {
	if header.Get("signature") == "" || header.Get("date") == "" {
		return errors.New("missing required field in header")
	}

	if err := inbox.verifySignature(ctx, header, payload, "/f/inbox"); err != nil {
		return fmt.Errorf("failed to verify signature: %w", err)
	}

	return inbox.ah.Dispatch(ctx, header, payload)
}

func (inbox *InboxUsecase) Validate(ctx context.Context, username string, header http.Header, payload InboxPayload) error {
	if header.Get("signature") == "" || header.Get("date") == "" {
		return errors.New("missing required field in header")
	}

	p, err := inbox.client.Ent.Profile.Query().
		Where(profile.Username(username), profile.DomainIsNil(), profile.StatusIsNil()).
		First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return fmt.Errorf("failed to get profile for validation: %w", err)
		}
		return nil
	}

	if err := inbox.verifySignature(ctx, header, payload, "/users/"+p.Username+"/inbox"); err != nil {
		return fmt.Errorf("failed to verify signature: %w", err)
	}

	return inbox.ah.Dispatch(ctx, header, payload)
}

func (inbox *InboxUsecase) verifySignature(ctx context.Context, header http.Header, payload InboxPayload, path string) error {
	date, err := parseDate(header)
	if err != nil {
		return fmt.Errorf("request expired: error=%w, date=%+v", err, date)
	}

	signature, err := parseSignature(header)
	if err != nil {
		return fmt.Errorf("invalid request: error=%w, signature=%s", err, header.Get("signature"))
	}

	if payload.Object.AttributedTo != nil {
		attr, err := url.Parse(*payload.Object.AttributedTo)
		if err != nil {
			return fmt.Errorf("failed to parse payload.Object.AttributedTo: error=%w, attributed_to=%s", err, *payload.Object.AttributedTo)
		}

		if attr.Host != signature.KeyId.Host {
			return fmt.Errorf("payload.Object.AttributedTo host mismatch: error=%w, attributed_to=%s, signature_key_id=%s", err, attr.Host, signature.KeyId.Host)
		}
	}

	objid, err := url.Parse(payload.Object.ID)
	if err != nil {
		return fmt.Errorf("failed to parse payload.Object.ID: error=%w, object_id=%s", err, payload.Object.ID)
	}
	if signature.KeyId.Host != objid.Host {
		return fmt.Errorf("payload.Object.ID host mismatch: object_id=%s, signature_key_id=%s", objid.Host, signature.KeyId.Host)
	}

	actor, err := inbox.client.Ent.Profile.Query().Where(profile.KeyID(signature.Raw["keyId"])).First(ctx)
	if err != nil {
		return fmt.Errorf("failed to get profile: %w", err)
	}

	pkey, err := loadPublicKey(actor.PublicKey)
	if err != nil {
		return fmt.Errorf("failed to load public key: error=%w", err)
	}

	if err := signature.Verify(pkey, header, payload, path); err != nil {
		return fmt.Errorf("failed to verify signature: error=%w", err)
	}

	return nil
}

func parseDate(h http.Header) (date time.Time, err error) {
	date, err = http.ParseTime(h.Get("date"))
	if err != nil {
		return
	}

	now := time.Now()
	if date.Before(now.AddDate(0, 0, -1)) || date.After(now.AddDate(0, 0, 1)) {
		return date, errors.New("date out of range")
	}

	return
}

func loadPublicKey(key string) (any, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, errors.New("failed to decode pem block")
	}

	return x509.ParsePKIXPublicKey(block.Bytes)
}

type signature struct {
	Raw map[string]string

	KeyId *url.URL
}

func parseSignature(h http.Header) (s signature, err error) {
	s.Raw = make(map[string]string)

	sigstr := h.Get("signature")
	parts := strings.Split(sigstr, ",")

	for _, p := range parts {
		key, value, ok := strings.Cut(p, "=")
		if !ok {
			continue
		}

		s.Raw[key] = strings.Trim(value, "\"")
	}

	if keystr, ok := s.Raw["keyId"]; !ok {
		return s, fmt.Errorf("missing keyId in signature header")
	} else if key, err := url.Parse(keystr); err != nil {
		return s, fmt.Errorf("invalid keyId format: %w", err)
	} else {
		s.KeyId = key
	}

	if _, ok := s.Raw["headers"]; !ok {
		return s, fmt.Errorf("missing headers in signature header")
	}
	if _, ok := s.Raw["signature"]; !ok {
		return s, fmt.Errorf("missing signature in signature header")
	}

	return
}

func (s *signature) Verify(pkey any, header http.Header, payload InboxPayload, path string) error {
	h := sha256.New()
	h.Write([]byte(payload.Raw))
	digest := "SHA-256=" + base64.StdEncoding.EncodeToString(h.Sum(nil))

	signedHeaders := strings.Split(s.Raw["headers"], " ")
	var lines []string
	for _, h := range signedHeaders {
		var value string
		switch h {
		case "(request-target)":
			value = "post " + path
		case "digest":
			value = digest
		default:
			value = header.Get(h)
		}
		lines = append(lines, h+": "+value)
	}
	signingString := strings.Join(lines, "\n")

	sigBytes, err := base64.StdEncoding.DecodeString(s.Raw["signature"])
	if err != nil {
		return fmt.Errorf("failed to decode signature: %w", err)
	}

	switch pub := pkey.(type) {
	case *rsa.PublicKey:
		hashed := sha256.Sum256([]byte(signingString))
		return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], sigBytes)
	case *ecdsa.PublicKey:
		hashed := sha256.Sum256([]byte(signingString))
		if !ecdsa.VerifyASN1(pub, hashed[:], sigBytes) {
			return errors.New("invalid ecdsa signature")
		}
		return nil
	case ed25519.PublicKey:
		if !ed25519.Verify(pub, []byte(signingString), sigBytes) {
			return errors.New("invalid ed25519 signature")
		}
		return nil
	default:
		return fmt.Errorf("unsupported public key type: %T", pkey)
	}
}
