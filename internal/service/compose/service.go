package compose

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	SearchLocation(w http.ResponseWriter, r *http.Request)
	ComposeSettings(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) SearchLocation(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Compose.SearchLocation")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ComposeSettings(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Compose.ComposeSettings")
	defer span.End()
	// TODO: Implement
}
