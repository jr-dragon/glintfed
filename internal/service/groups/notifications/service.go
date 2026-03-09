package notifications

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	SelfGlobalNotifications(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) SelfGlobalNotifications(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsNotifications.SelfGlobalNotifications")
	defer span.End()
	// TODO: Implement
}
