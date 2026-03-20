package worker

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"glintfed.org/ent"
)

func generateTestRSAKey(t *testing.T) (*rsa.PrivateKey, string) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pubASN1, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	require.NoError(t, err)

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return key, string(pubPEM)
}

func signRequest(t *testing.T, priv *rsa.PrivateKey, method, path, date string, body string, keyID string) string {
	h := sha256.New()
	h.Write([]byte(body))
	digest := base64.StdEncoding.EncodeToString(h.Sum(nil))
	digestHeader := "SHA-256=" + digest

	signingString := fmt.Sprintf("(request-target): %s %s\ndate: %s\ndigest: %s", method, path, date, digestHeader)

	hashed := sha256.Sum256([]byte(signingString))
	sig, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
	require.NoError(t, err)

	sigBase64 := base64.StdEncoding.EncodeToString(sig)

	return fmt.Sprintf(`keyId="%s",algorithm="rsa-sha256",headers="(request-target) date digest",signature="%s"`, keyID, sigBase64)
}

func TestInboxUsecase_Delete(t *testing.T) {
	t.Run("missing header", func(t *testing.T) {
		inbox := &InboxUsecase{}
		err := inbox.Delete(context.Background(), http.Header{}, InboxPayload{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing required field in header")
	})

	t.Run("valid delete person activity", func(t *testing.T) {
		priv, pubStr := generateTestRSAKey(t)
		keyID := "https://example.com/actor#main-key"
		actorID := "https://example.com/actor"
		date := time.Now().UTC().Format(http.TimeFormat)
		path := "/f/inbox"
		payload := InboxPayload{
			Raw:   `{"type":"Delete","actor":"https://example.com/actor","object":{"id":"https://example.com/actor","type":"Person"}}`,
			Type:  new("Delete"),
			Actor: new(actorID),
			Object: &InboxPayloadObject{
				ID:   actorID,
				Type: "Person",
			},
		}

		sig := signRequest(t, priv, "post", path, date, payload.Raw, keyID)

		header := http.Header{}
		header.Set("Date", date)
		header.Set("Signature", sig)

		pg := &ProfileGetterMock{
			GetByKeyIDFunc: func(ctx context.Context, kid string) (*ent.Profile, error) {
				assert.Equal(t, keyID, kid)
				return &ent.Profile{
					PublicKey: pubStr,
				}, nil
			},
			GetActiveRemoteProfileFunc: func(ctx context.Context, url string) (*ent.Profile, error) {
				assert.Equal(t, actorID, url)
				return &ent.Profile{ID: 123}, nil
			},
		}

		pr := &ProfileRemoverMock{
			RemoteProfileFunc: func(ctx context.Context, profile *ent.Profile) error {
				assert.Equal(t, uint64(123), profile.ID)
				return nil
			},
		}

		inbox := NewInboxUsecase(nil, pg, pr, nil)
		err := inbox.Delete(context.Background(), header, payload)
		assert.NoError(t, err)
		assert.Len(t, pr.RemoteProfileCalls(), 1)
	})

	t.Run("valid delete other activity (dispatch to ActivityDispatcher)", func(t *testing.T) {
		priv, pubStr := generateTestRSAKey(t)
		keyID := "https://example.com/actor#main-key"
		actorID := "https://example.com/actor"
		date := time.Now().UTC().Format(http.TimeFormat)
		path := "/f/inbox"
		payload := InboxPayload{
			Raw:   `{"type":"Delete","actor":"https://example.com/actor","object":{"id":"https://example.com/note/123","type":"Note"}}`,
			Type:  new("Delete"),
			Actor: new(actorID),
			Object: &InboxPayloadObject{
				ID:   "https://example.com/note/123",
				Type: "Note",
			},
		}

		sig := signRequest(t, priv, "post", path, date, payload.Raw, keyID)

		header := http.Header{}
		header.Set("Date", date)
		header.Set("Signature", sig)

		pg := &ProfileGetterMock{
			GetByKeyIDFunc: func(ctx context.Context, kid string) (*ent.Profile, error) {
				return &ent.Profile{
					PublicKey: pubStr,
				}, nil
			},
		}

		ad := &ActivityDispatcherMock{
			DeleteFunc: func(ctx context.Context, header http.Header, p InboxPayload) error {
				assert.Equal(t, payload.Object.ID, p.Object.ID)
				return nil
			},
		}

		inbox := NewInboxUsecase(nil, pg, nil, ad)
		err := inbox.Delete(context.Background(), header, payload)
		assert.NoError(t, err)
		assert.Len(t, ad.DeleteCalls(), 1)
	})
}

func TestInboxUsecase_Inbox(t *testing.T) {
	t.Run("missing header", func(t *testing.T) {
		inbox := &InboxUsecase{}
		err := inbox.Inbox(context.Background(), http.Header{}, InboxPayload{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing required field in header")
	})

	t.Run("valid dispatch", func(t *testing.T) {
		priv, pubStr := generateTestRSAKey(t)
		keyID := "https://example.com/actor#main-key"
		actorID := "https://example.com/actor"
		date := time.Now().UTC().Format(http.TimeFormat)
		path := "/f/inbox"
		payload := InboxPayload{
			Raw:   `{"type":"Create","actor":"https://example.com/actor","object":{"id":"https://example.com/note/123","type":"Note"}}`,
			Type:  new("Create"),
			Actor: new(actorID),
			Object: &InboxPayloadObject{
				ID:   "https://example.com/note/123",
				Type: "Note",
			},
		}

		sig := signRequest(t, priv, "post", path, date, payload.Raw, keyID)

		header := http.Header{}
		header.Set("Date", date)
		header.Set("Signature", sig)

		pg := &ProfileGetterMock{
			GetByKeyIDFunc: func(ctx context.Context, kid string) (*ent.Profile, error) {
				return &ent.Profile{
					PublicKey: pubStr,
				}, nil
			},
		}

		ad := &ActivityDispatcherMock{
			DispatchFunc: func(ctx context.Context, h http.Header, p InboxPayload) error {
				assert.Equal(t, payload.Object.ID, p.Object.ID)
				return nil
			},
		}

		inbox := NewInboxUsecase(nil, pg, nil, ad)
		err := inbox.Inbox(context.Background(), header, payload)
		assert.NoError(t, err)
		assert.Len(t, ad.DispatchCalls(), 1)
	})
}

func TestInboxUsecase_Validate(t *testing.T) {
	t.Run("missing header", func(t *testing.T) {
		inbox := &InboxUsecase{}
		err := inbox.Validate(context.Background(), "user", http.Header{}, InboxPayload{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing required field in header")
	})

	t.Run("profile not found", func(t *testing.T) {
		pg := &ProfileGetterMock{
			GetActiveLocalProfileFunc: func(ctx context.Context, username string) (*ent.Profile, error) {
				return nil, &ent.NotFoundError{}
			},
		}
		inbox := &InboxUsecase{pg: pg}
		header := http.Header{}
		header.Set("Signature", "sig")
		header.Set("Date", time.Now().UTC().Format(http.TimeFormat))

		err := inbox.Validate(context.Background(), "nonexistent", header, InboxPayload{})
		assert.NoError(t, err)
	})

	t.Run("valid validation and dispatch", func(t *testing.T) {
		priv, pubStr := generateTestRSAKey(t)
		username := "testuser"
		keyID := "https://example.com/actor#main-key"
		actorID := "https://example.com/actor"
		date := time.Now().UTC().Format(http.TimeFormat)
		path := "/users/" + username + "/inbox"
		payload := InboxPayload{
			Raw:   `{"type":"Follow","actor":"https://example.com/actor","object":{"id":"https://example.com/note/123","type":"Note"}}`,
			Type:  new("Follow"),
			Actor: new(actorID),
			Object: &InboxPayloadObject{
				ID:   "https://example.com/note/123",
				Type: "Note",
			},
		}

		sig := signRequest(t, priv, "post", path, date, payload.Raw, keyID)

		header := http.Header{}
		header.Set("Date", date)
		header.Set("Signature", sig)

		pg := &ProfileGetterMock{
			GetActiveLocalProfileFunc: func(ctx context.Context, u string) (*ent.Profile, error) {
				assert.Equal(t, username, u)
				return &ent.Profile{Username: username}, nil
			},
			GetByKeyIDFunc: func(ctx context.Context, kid string) (*ent.Profile, error) {
				return &ent.Profile{
					PublicKey: pubStr,
				}, nil
			},
		}

		ad := &ActivityDispatcherMock{
			DispatchFunc: func(ctx context.Context, h http.Header, p InboxPayload) error {
				return nil
			},
		}

		inbox := NewInboxUsecase(nil, pg, nil, ad)
		err := inbox.Validate(context.Background(), username, header, payload)
		assert.NoError(t, err)
		assert.Len(t, ad.DispatchCalls(), 1)
	})
}
