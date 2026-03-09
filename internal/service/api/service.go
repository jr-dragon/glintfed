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

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "Api."+name)
	defer span.End()
}

func (s *svc) AvatarUpdate(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "AvatarUpdate") }
func (s *svc) Notifications(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Notifications") }
func (s *svc) VerifyCredentials(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "VerifyCredentials")
}
func (s *svc) AccountLikes(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "AccountLikes") }
func (s *svc) Archive(w http.ResponseWriter, r *http.Request)       { s.stub(w, r, "Archive") }
func (s *svc) Unarchive(w http.ResponseWriter, r *http.Request)     { s.stub(w, r, "Unarchive") }
func (s *svc) ArchivedPosts(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "ArchivedPosts") }
func (s *svc) SiteConfiguration(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "SiteConfiguration")
}
func (s *svc) UserRecommendations(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "UserRecommendations")
}
