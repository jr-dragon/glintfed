package topic

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GroupTopics(w http.ResponseWriter, r *http.Request)
	GroupTopicTag(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) GroupTopics(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsTopic.GroupTopics")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GroupTopicTag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsTopic.GroupTopicTag")
	defer span.End()
	// TODO: Implement
}
