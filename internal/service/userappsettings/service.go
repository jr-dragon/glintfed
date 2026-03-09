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

func (s *svc) Get(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "UserAppSettings.Get")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "UserAppSettings.Store")
	defer span.End()
	// TODO: Implement
}
