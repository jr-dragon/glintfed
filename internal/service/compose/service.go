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

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "Compose."+name)
	defer span.End()
}

func (s *svc) SearchLocation(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "SearchLocation") }
func (s *svc) ComposeSettings(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "ComposeSettings")
}
