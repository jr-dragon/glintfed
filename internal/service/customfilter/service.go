package customfilter

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) Index(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Index")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Show(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Show")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Store")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Update(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Update")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Delete")
	defer span.End()
	// TODO: Implement
}
