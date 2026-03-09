package domainblock

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Index(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "DomainBlock."+name)
	defer span.End()
}

func (s *svc) Index(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "Index") }
func (s *svc) Store(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "Store") }
func (s *svc) Delete(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Delete") }
