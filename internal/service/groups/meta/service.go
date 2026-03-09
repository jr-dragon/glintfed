package meta

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	DeleteGroup(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMeta.DeleteGroup")
	defer span.End()
	// TODO: Implement
}
