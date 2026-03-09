package healthcheck

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) Get(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "HealthCheck.Get")
	defer span.End()

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "max-age=0, must-revalidate, no-cache, no-store")
	w.Write([]byte("OK"))
}
