package search

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	InviteFriendsToGroup(w http.ResponseWriter, r *http.Request)
	SearchFriendsToInvite(w http.ResponseWriter, r *http.Request)
	SearchGlobalResults(w http.ResponseWriter, r *http.Request)
	SearchLocalAutocomplete(w http.ResponseWriter, r *http.Request)
	SearchAddRecent(w http.ResponseWriter, r *http.Request)
	SearchGetRecent(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) InviteFriendsToGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.InviteFriendsToGroup")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SearchFriendsToInvite(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchFriendsToInvite")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SearchGlobalResults(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchGlobalResults")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SearchLocalAutocomplete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchLocalAutocomplete")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SearchAddRecent(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchAddRecent")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SearchGetRecent(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchGetRecent")
	defer span.End()
	// TODO: Implement
}
