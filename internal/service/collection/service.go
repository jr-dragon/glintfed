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

func (s *svc) GetUserCollections(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.AvatarUpdate")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetItems(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.GetItems")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetCollection(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.GetCollection")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StoreId(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.StoreId")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.Store")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.Delete")
	defer span.End()
	// TODO: Implement
}

func (s *svc) DeleteId(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.DeleteId")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetSelfCollections(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.GetSelfCollections")
	defer span.End()
	// TODO: Implement
}
