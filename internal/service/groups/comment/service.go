package comment

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetComments(w http.ResponseWriter, r *http.Request)
	StoreComment(w http.ResponseWriter, r *http.Request)
	StoreCommentPhoto(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)
	LikePost(w http.ResponseWriter, r *http.Request)
	UnlikePost(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) GetComments(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.GetComments")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StoreComment(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.StoreComment")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StoreCommentPhoto(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.StoreCommentPhoto")
	defer span.End()
	// TODO: Implement
}

func (s *svc) DeleteComment(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.DeleteComment")
	defer span.End()
	// TODO: Implement
}

func (s *svc) LikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.LikePost")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UnlikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.UnlikePost")
	defer span.End()
	// TODO: Implement
}
