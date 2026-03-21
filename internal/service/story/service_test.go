package story

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

func TestService_GetActivityObject(t *testing.T) {
	cfg := &data.Config{
		App: data.AppConfig{
			Url: "https://example.com",
			Instance: data.InstanceConfig{
				Stories: data.StoriesConfig{
					Enabled: true,
				},
			},
		},
	}

	now := time.Now()
	st := &ent.Story{
		ID:           123,
		BearcapToken: "valid-token",
		ExpiresAt:    now.Add(24 * time.Hour),
		CreatedAt:    now.Add(-5 * time.Minute),
		Type:         "photo",
		Mime:         "image/jpeg",
		Duration:     15,
		CanReply:     true,
		CanReact:     true,
	}
	// Setup profile edge
	st.Edges.Profile = &ent.Profile{
		Username: "testuser",
	}

	t.Run("success", func(t *testing.T) {
		sg := &StoryGetterMock{
			GetByUsernameAndIDFunc: func(ctx context.Context, username string, id uint64) (*ent.Story, error) {
				assert.Equal(t, "testuser", username)
				assert.Equal(t, uint64(123), id)
				return st, nil
			},
		}
		svc := New(cfg, sg)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/stories/testuser/123", nil)
		r.Header.Set("Accept", "application/json")
		r.Header.Set("Authorization", "Bearer valid-token")

		// Add chi context for URL params
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("username", "testuser")
		rctx.URLParams.Add("id", "123")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		svc.GetActivityObject(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/activity+json", w.Header().Get("Content-Type"))

		var resp GetActivityObjectResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, "Story", resp.Type)
		assert.Equal(t, "Image", resp.Attachment.Type)
		assert.Equal(t, uint16(15), resp.Duration)
	})

	t.Run("stories disabled", func(t *testing.T) {
		disabledCfg := *cfg
		disabledCfg.App.Instance.Stories.Enabled = false

		svc := New(&disabledCfg, &StoryGetterMock{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/stories/testuser/123", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("username", "testuser")
		rctx.URLParams.Add("id", "123")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		svc.GetActivityObject(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("redirect if not json", func(t *testing.T) {
		svc := New(cfg, &StoryGetterMock{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/stories/testuser/123", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("username", "testuser")
		rctx.URLParams.Add("id", "123")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		svc.GetActivityObject(w, r)
		assert.Equal(t, http.StatusFound, w.Code)
		assert.Equal(t, "/stories/testuser", w.Header().Get("Location"))
	})

	t.Run("invalid token", func(t *testing.T) {
		sg := &StoryGetterMock{
			GetByUsernameAndIDFunc: func(ctx context.Context, username string, id uint64) (*ent.Story, error) {
				return st, nil
			},
		}
		svc := New(cfg, sg)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/stories/testuser/123", nil)
		r.Header.Set("Accept", "application/json")
		r.Header.Set("Authorization", "Bearer invalid-token")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("username", "testuser")
		rctx.URLParams.Add("id", "123")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		svc.GetActivityObject(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("expired story", func(t *testing.T) {
		expiredSt := *st
		expiredSt.ExpiresAt = now.Add(-1 * time.Hour)

		sg := &StoryGetterMock{
			GetByUsernameAndIDFunc: func(ctx context.Context, username string, id uint64) (*ent.Story, error) {
				return &expiredSt, nil
			},
		}
		svc := New(cfg, sg)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/stories/testuser/123", nil)
		r.Header.Set("Accept", "application/json")
		r.Header.Set("Authorization", "Bearer valid-token")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("username", "testuser")
		rctx.URLParams.Add("id", "123")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		svc.GetActivityObject(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
