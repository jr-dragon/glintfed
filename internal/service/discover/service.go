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

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "Discover."+name)
	defer span.End()
}

func (s *svc) TrendingApi(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "TrendingApi") }
func (s *svc) TrendingHashtags(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "TrendingHashtags")
}
func (s *svc) DiscoverNetworkTrending(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "DiscoverNetworkTrending")
}
