package api

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetConfig(w http.ResponseWriter, r *http.Request)
	GetGroupAccount(w http.ResponseWriter, r *http.Request)
	GetGroupCategories(w http.ResponseWriter, r *http.Request)
	GetGroupsByCategory(w http.ResponseWriter, r *http.Request)
	GetRecommendedGroups(w http.ResponseWriter, r *http.Request)
	GetSelfGroups(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) GetConfig(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetConfig")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupAccount(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetGroupAccount")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupCategories(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetGroupCategories")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupsByCategory(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetGroupsByCategory")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetRecommendedGroups(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetRecommendedGroups")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetSelfGroups(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetSelfGroups")
	defer span.End()
	// TODO: Implement
}
