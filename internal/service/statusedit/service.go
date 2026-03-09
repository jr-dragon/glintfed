package statusedit

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Store(w http.ResponseWriter, r *http.Request)
	History(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "StatusEdit.Store")
	defer span.End()
	// TODO: Implement
}

func (s *svc) History(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "StatusEdit.History")
	defer span.End()
	// TODO: Implement
}
