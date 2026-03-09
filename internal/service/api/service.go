package api

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	AvatarUpdate(w http.ResponseWriter, r *http.Request)
	Notifications(w http.ResponseWriter, r *http.Request)
	VerifyCredentials(w http.ResponseWriter, r *http.Request)
	AccountLikes(w http.ResponseWriter, r *http.Request)
	Archive(w http.ResponseWriter, r *http.Request)
	Unarchive(w http.ResponseWriter, r *http.Request)
	ArchivedPosts(w http.ResponseWriter, r *http.Request)
	SiteConfiguration(w http.ResponseWriter, r *http.Request)
	UserRecommendations(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) AvatarUpdate(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AvatarUpdate")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Notifications(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.Notifications")
	defer span.End()
	// TODO: Implement
}

func (s *svc) VerifyCredentials(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.VerifyCredentials")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountLikes(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AccountLikes")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Archive(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.Archive")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Unarchive(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.Unarchive")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ArchivedPosts(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ArchivedPosts")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SiteConfiguration(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.SiteConfiguration")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UserRecommendations(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.UserRecommendations")
	defer span.End()
	// TODO: Implement
}
