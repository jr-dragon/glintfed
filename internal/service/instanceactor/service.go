package instanceactor

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Profile(w http.ResponseWriter, r *http.Request)
	Inbox(w http.ResponseWriter, r *http.Request)
	Outbox(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) Profile(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Profile")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Inbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Inbox")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Outbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Outbox")
	defer span.End()
	// TODO: Implement
}
