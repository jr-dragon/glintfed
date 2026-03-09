package discover

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetDiscoverPopular(w http.ResponseWriter, r *http.Request)
	GetDiscoverNew(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) GetDiscoverPopular(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsDiscover.GetDiscoverPopular")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetDiscoverNew(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsDiscover.GetDiscoverNew")
	defer span.End()
	// TODO: Implement
}
