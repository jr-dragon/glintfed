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

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "StatusEdit."+name)
	defer span.End()
}

func (s *svc) Store(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "Store") }
func (s *svc) History(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "History") }
