package apiv2

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Instance(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	GetWebsocketConfig(w http.ResponseWriter, r *http.Request)
	MediaUploadV2(w http.ResponseWriter, r *http.Request)
	StatusContextV2(w http.ResponseWriter, r *http.Request)
	StatusDescendants(w http.ResponseWriter, r *http.Request)
	StatusAncestors(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "ApiV2."+name)
	defer span.End()
}

func (s *svc) Instance(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "Instance") }
func (s *svc) Search(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "Search") }
func (s *svc) GetWebsocketConfig(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "GetWebsocketConfig")
}
func (s *svc) MediaUploadV2(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "MediaUploadV2") }
func (s *svc) StatusContextV2(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "StatusContextV2")
}
func (s *svc) StatusDescendants(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "StatusDescendants")
}
func (s *svc) StatusAncestors(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "StatusAncestors")
}
