package landing

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetDirectoryApi(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "Landing."+name)
	defer span.End()
}

func (s *svc) GetDirectoryApi(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "GetDirectoryApi")
}
