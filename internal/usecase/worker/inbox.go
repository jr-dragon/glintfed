package worker

import (
	"context"
	"crypto/x509"
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

type InboxUsecase struct {
	client *data.Client
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
	}
}

func (inbox *InboxUsecase) Delete(ctx context.Context, header http.Header, payload InboxPayload) {
	if header.Get("signature") == "" || header.Get("date") == "" {
		slog.ErrorContext(ctx, "missing required field in header", slog.Any("header", header))
		return
	}

	if err := inbox.verifySignature(ctx, header, payload); err != nil {
		slog.ErrorContext(ctx, "failed to verify signature", logs.ErrAttr(err))
		return
	}

	if payload.Object.Type == "Person" && payload.Actor != nil && *payload.Actor == payload.Object.ID {
		// todo
	} else {
		// todo
	}
}

func (inbox *InboxUsecase) Inbox(ctx context.Context, header http.Header, payload InboxPayload) {

}

func (inbox *InboxUsecase) Validate(ctx context.Context, username string, header http.Header, payload InboxPayload) {

}

func (inbox *InboxUsecase) verifySignature(ctx context.Context, header http.Header, payload InboxPayload) error {
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

	actor, err := inbox.client.Ent.Profile.Query().Where(profile.KeyIDEQ(signature.Raw["keyId"])).First(ctx)
	if ent.IsNotFound(err) {
		// todo: first or create if remote url, failed if local url
	} else if err != nil {
		return fmt.Errorf("failed to get profile: %w", err)
	}

	pkey, err := loadPublicKey(actor.PublicKey)
	if err != nil {
		return fmt.Errorf("failed to load public key: error=%w", err)
	}

	if err := signature.Verify(pkey, header, payload, "/f/inbox"); err != nil {
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

		s.Raw[key] = value
	}

	if keystr, ok := s.Raw["keyId"]; !ok {
		return s, fmt.Errorf("missing keyId in signature header")
	} else if key, err := url.Parse(keystr); err != nil {
		s.KeyId = key
		return s, fmt.Errorf("invalid keyId format: %w", err)
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
	// TODO
	return nil
}
