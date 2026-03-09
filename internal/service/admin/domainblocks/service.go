package domainblocks

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "AdminDomainBlocks."+name)
	defer span.End()
}

func (s *svc) Index(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "Index") }
func (s *svc) Show(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "Show") }
func (s *svc) Create(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Create") }
func (s *svc) Update(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Update") }
func (s *svc) Delete(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Delete") }
