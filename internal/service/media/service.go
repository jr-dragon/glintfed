package media

import (
	"context"
	"log/slog"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/lib/urls"
	"glintfed.org/internal/service/internal"
)

type Service interface {
	FallbackRedirect(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_media_getter.go . MediaGetter
type MediaGetter interface {
	GetCDNUrl(ctx context.Context, path string) (string, error)
}

func New(cfg *data.Config, mg MediaGetter) Service {
	return &svc{
		cfg: cfg,
		mg:  mg,
	}
}

type svc struct {
	cfg *data.Config

	mg MediaGetter
}

func (s *svc) FallbackRedirect(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Media.FallbackRedirect")
	defer span.End()

	if !s.cfg.App.CloudStorage {
		http.Redirect(w, r, urls.MustJoinPath(s.cfg.App.Url, "storage/no-preview.png"), http.StatusFound)
		return
	}

	url, err := s.mg.GetCDNUrl(
		r.Context(),
		path.Join("public/m/v2", chi.URLParam(r, "pid"), chi.URLParam(r, "mhash"), chi.URLParam(r, "uhash"), chi.URLParam(r, "f")),
	)
	if err != nil {
		if ent.IsNotFound(err) {
			http.Redirect(w, r, urls.MustJoinPath(s.cfg.App.Url, "storage/no-preview.png"), http.StatusFound)
		} else {
			slog.ErrorContext(r.Context(), "failed to get media", logs.ErrAttr(err))
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
