package post

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	StorePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
	LikePost(w http.ResponseWriter, r *http.Request)
	UnlikePost(w http.ResponseWriter, r *http.Request)
	GetGroupMedia(w http.ResponseWriter, r *http.Request)
	GetStatus(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) StorePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.StorePost")
	defer span.End()
	// TODO: Implement
}

func (s *svc) DeletePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.DeletePost")
	defer span.End()
	// TODO: Implement
}

func (s *svc) LikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.LikePost")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UnlikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.UnlikePost")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupMedia(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupPost.GetGroupMedia")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetStatus(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.GetStatus")
	defer span.End()
	// TODO: Implement
}
