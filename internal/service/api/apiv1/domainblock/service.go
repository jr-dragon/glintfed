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

func (s *svc) Index(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DomainBlocks.Index")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DomainBlocks.Store")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DomainBlocks.Delete")
	defer span.End()
	// TODO: Implement
}
