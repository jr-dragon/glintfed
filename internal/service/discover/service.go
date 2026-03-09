package discover

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	TrendingApi(w http.ResponseWriter, r *http.Request)
	TrendingHashtags(w http.ResponseWriter, r *http.Request)
	DiscoverNetworkTrending(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) TrendingApi(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Discover.TrendingApi")
	defer span.End()
	// TODO: Implement
}

func (s *svc) TrendingHashtags(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Discover.TrendingHashtags")
	defer span.End()
	// TODO: Implement
}
func (s *svc) DiscoverNetworkTrending(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Discover.DiscoverNetworkTrending")
	defer span.End()
	// TODO: Implement
}
