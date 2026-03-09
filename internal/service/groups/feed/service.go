package feed

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetSelfFeed(w http.ResponseWriter, r *http.Request)
	GetGroupProfileFeed(w http.ResponseWriter, r *http.Request)
	GetGroupFeed(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) GetSelfFeed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsFeed.GetSelfFeed")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupProfileFeed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsFeed.GetGroupProfileFeed")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupFeed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsFeed.GetGroupFeed")
	defer span.End()
	// TODO: Implement
}
