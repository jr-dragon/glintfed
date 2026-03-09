package tags

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	RelatedTags(w http.ResponseWriter, r *http.Request)
	FollowHashtag(w http.ResponseWriter, r *http.Request)
	UnfollowHashtag(w http.ResponseWriter, r *http.Request)
	GetHashtag(w http.ResponseWriter, r *http.Request)
	GetFollowedTags(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "Tags."+name)
	defer span.End()
}

func (s *svc) RelatedTags(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "RelatedTags") }
func (s *svc) FollowHashtag(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "FollowHashtag") }
func (s *svc) UnfollowHashtag(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "UnfollowHashtag")
}
func (s *svc) GetHashtag(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetHashtag") }
func (s *svc) GetFollowedTags(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "GetFollowedTags")
}
