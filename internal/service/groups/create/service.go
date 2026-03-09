package create

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	CheckCreatePermission(w http.ResponseWriter, r *http.Request)
	StoreGroup(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) CheckCreatePermission(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsCreate.CheckCreatePermission")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StoreGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsCreate.StoreGroup")
	defer span.End()
	// TODO: Implement
}
