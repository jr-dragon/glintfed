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

func (s *svc) Instance(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.Instance")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Search(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.Search")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetWebsocketConfig(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.GetWebsocketConfig")
	defer span.End()
	// TODO: Implement
}

func (s *svc) MediaUploadV2(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.MediaUploadV2")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusContextV2(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.StatusContextV2")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusDescendants(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.StatusDescendants")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusAncestors(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.StatusAncestors")
	defer span.End()
	// TODO: Implement
}
