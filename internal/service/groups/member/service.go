package member

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetGroupMembers(w http.ResponseWriter, r *http.Request)
	GetGroupMemberJoinRequests(w http.ResponseWriter, r *http.Request)
	HandleGroupMemberJoinRequest(w http.ResponseWriter, r *http.Request)
	GetGroupMember(w http.ResponseWriter, r *http.Request)
	GetGroupMemberCommonIntersections(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.GetGroupMembers")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupMemberJoinRequests(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.GetGroupMemberJoinRequests")
	defer span.End()
	// TODO: Implement
}

func (s *svc) HandleGroupMemberJoinRequest(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.HandleGroupMemberJoinRequest")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupMember(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.GetGroupMember")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetGroupMemberCommonIntersections(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.GetGroupMemberCommonIntersections")
	defer span.End()
	// TODO: Implement
}
