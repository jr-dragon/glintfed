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

func (s *svc) RelatedTags(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.RelatedTags")
	defer span.End()
	// TODO: Implement
}

func (s *svc) FollowHashtag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.FollowHashtag")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UnfollowHashtag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.UnfollowHashtag")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetHashtag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.GetHashtag")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetFollowedTags(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.GetFollowedTags")
	defer span.End()
	// TODO: Implement
}
