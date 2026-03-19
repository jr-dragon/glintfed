package federation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

func TestWebfinger(t *testing.T) {
	cfg := &data.Config{
		App: data.AppConfig{
			Url:           "https://glintfed.test",
			MaxNameLength: 20,
			Federation: data.FederationConfig{
				Webfinger: data.WebfingerConfig{
					Enabled: true,
				},
				Activitypub: data.ActivitypubConfig{
					SharedInbox: true,
				},
			},
		},
	}

	t.Run("bad_request_if_disabled", func(t *testing.T) {
		disabledCfg := *cfg
		disabledCfg.App.Federation.Webfinger.Enabled = false
		s := New(&disabledCfg, nil, nil, nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource=acct:alice@glintfed.test", nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("bad_request_if_missing_resource", func(t *testing.T) {
		s := New(cfg, nil, nil, nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger", nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("shared_inbox_resource", func(t *testing.T) {
		s := New(cfg, nil, nil, nil, nil)

		resource := "acct:glintfed.test@glintfed.test"
		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource="+resource, nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp SharedInboxResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, resource, resp.Subject)
		assert.Contains(t, resp.Aliases, "https://glintfed.test/i/actor")
	})

	t.Run("user_resource_success", func(t *testing.T) {
		puc := &ProfileUsecaseMock{
			GetByUsernameFunc: func(ctx context.Context, username string) (*ent.Profile, error) {
				if username == "alice" {
					return &ent.Profile{Username: "alice", AvatarURL: "https://glintfed.test/avatar.webp"}, nil
				}
				return nil, fmt.Errorf("not found")
			},
		}
		s := New(cfg, nil, puc, nil, nil)

		resource := "https://glintfed.test/users/alice"
		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource="+resource, nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp webfinger
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, "acct:alice@glintfed.test", resp.Subject)
		assert.Contains(t, resp.Aliases, "https://glintfed.test/alice")
	})

	t.Run("user_resource_acct_success", func(t *testing.T) {
		puc := &ProfileUsecaseMock{
			GetByUsernameFunc: func(ctx context.Context, username string) (*ent.Profile, error) {
				if username == "alice" {
					return &ent.Profile{Username: "alice", AvatarURL: "https://glintfed.test/avatar.webp"}, nil
				}
				return nil, fmt.Errorf("not found")
			},
		}
		s := New(cfg, nil, puc, nil, nil)

		resource := "acct:alice@glintfed.test"
		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource="+resource, nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp webfinger
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, "acct:alice@glintfed.test", resp.Subject)
		assert.Contains(t, resp.Aliases, "https://glintfed.test/alice")
	})

	t.Run("user_resource_not_found", func(t *testing.T) {
		puc := &ProfileUsecaseMock{
			GetByUsernameFunc: func(ctx context.Context, username string) (*ent.Profile, error) {
				return nil, fmt.Errorf("not found")
			},
		}
		s := New(cfg, nil, puc, nil, nil)

		resource := "https://glintfed.test/users/bob"
		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource="+resource, nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("username_too_long", func(t *testing.T) {
		s := New(cfg, nil, nil, nil, nil)

		resource := "https://glintfed.test/users/thisusernameiswaytoolongtobevalid"
		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource="+resource, nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid_username_characters", func(t *testing.T) {
		s := New(cfg, nil, nil, nil, nil)

		resource := "https://glintfed.test/users/alice!"
		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource="+resource, nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid_resource_domain", func(t *testing.T) {
		s := New(cfg, nil, nil, nil, nil)

		resource := "https://otherdomain.test/users/alice"
		req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource="+resource, nil)
		w := httptest.NewRecorder()

		s.Webfinger(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
