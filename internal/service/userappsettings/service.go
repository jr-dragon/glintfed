package userappsettings

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Get(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "UserAppSettings."+name)
	defer span.End()
}

func (s *svc) Get(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "Get") }
func (s *svc) Store(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Store") }
