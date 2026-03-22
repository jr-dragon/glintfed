package media

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

func TestService_FallbackRedirect(t *testing.T) {
	cfg := &data.Config{
		App: data.AppConfig{
			Url:          "https://example.com",
			CloudStorage: true,
		},
	}

	t.Run("success", func(t *testing.T) {
		mg := &MediaGetterMock{
			GetCDNUrlFunc: func(ctx context.Context, path string) (string, error) {
				assert.Equal(t, "public/m/v2/pid/mhash/uhash/f.jpg", path)
				return "https://cdn.example.com/media.jpg", nil
			},
		}
		svc := New(cfg, mg)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/m/pid/mhash/uhash/f.jpg", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("pid", "pid")
		rctx.URLParams.Add("mhash", "mhash")
		rctx.URLParams.Add("uhash", "uhash")
		rctx.URLParams.Add("f", "f.jpg")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		svc.FallbackRedirect(w, r)

		assert.Equal(t, http.StatusFound, w.Code)
		assert.Equal(t, "https://cdn.example.com/media.jpg", w.Header().Get("Location"))
	})

	t.Run("cloud storage disabled", func(t *testing.T) {
		disabledCfg := *cfg
		disabledCfg.App.CloudStorage = false

		svc := New(&disabledCfg, &MediaGetterMock{})

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/m/pid/mhash/uhash/f.jpg", nil)

		svc.FallbackRedirect(w, r)

		assert.Equal(t, http.StatusFound, w.Code)
		assert.Equal(t, "https://example.com/storage/no-preview.png", w.Header().Get("Location"))
	})

	t.Run("not found", func(t *testing.T) {
		mg := &MediaGetterMock{
			GetCDNUrlFunc: func(ctx context.Context, path string) (string, error) {
				return "", &ent.NotFoundError{}
			},
		}
		svc := New(cfg, mg)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/m/pid/mhash/uhash/f.jpg", nil)

		svc.FallbackRedirect(w, r)

		assert.Equal(t, http.StatusFound, w.Code)
		assert.Equal(t, "https://example.com/storage/no-preview.png", w.Header().Get("Location"))
	})

	t.Run("internal error", func(t *testing.T) {
		mg := &MediaGetterMock{
			GetCDNUrlFunc: func(ctx context.Context, path string) (string, error) {
				return "", errors.New("db error")
			},
		}
		svc := New(cfg, mg)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/m/pid/mhash/uhash/f.jpg", nil)

		svc.FallbackRedirect(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
