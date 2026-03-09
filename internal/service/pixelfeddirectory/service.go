package pixelfeddirectory

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
	_, span := internal.T.Start(r.Context(), "PixelfedDirectory.Get")
	defer span.End()
	// TODO: Implement
}
