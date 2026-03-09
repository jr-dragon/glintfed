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

func (s *svc) Index(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Index")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Show(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Show")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Create(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Create")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Update(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Update")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Delete")
	defer span.End()
	// TODO: Implement
}
