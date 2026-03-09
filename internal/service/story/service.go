package story

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetActivityObject(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) GetActivityObject(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Story.GetActivityObject")
	defer span.End()
	// TODO: Implement
}
