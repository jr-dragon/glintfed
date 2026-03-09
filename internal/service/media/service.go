package media

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	FallbackRedirect(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) FallbackRedirect(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Media.FallbackRedirect")
	defer span.End()
	// TODO: Implement
}
