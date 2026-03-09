package collection

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetUserCollections(w http.ResponseWriter, r *http.Request)
	GetItems(w http.ResponseWriter, r *http.Request)
	GetCollection(w http.ResponseWriter, r *http.Request)
	StoreId(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	DeleteId(w http.ResponseWriter, r *http.Request)
	GetSelfCollections(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "Collection."+name)
	defer span.End()
}

func (s *svc) GetUserCollections(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "GetUserCollections")
}
func (s *svc) GetItems(w http.ResponseWriter, r *http.Request)      { s.stub(w, r, "GetItems") }
func (s *svc) GetCollection(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetCollection") }
func (s *svc) StoreId(w http.ResponseWriter, r *http.Request)       { s.stub(w, r, "StoreId") }
func (s *svc) Store(w http.ResponseWriter, r *http.Request)         { s.stub(w, r, "Store") }
func (s *svc) Delete(w http.ResponseWriter, r *http.Request)        { s.stub(w, r, "Delete") }
func (s *svc) DeleteId(w http.ResponseWriter, r *http.Request)      { s.stub(w, r, "DeleteId") }
func (s *svc) GetSelfCollections(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "GetSelfCollections")
}
