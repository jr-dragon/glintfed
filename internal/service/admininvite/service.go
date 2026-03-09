package admininvite

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	ApiVerifyCheck(w http.ResponseWriter, r *http.Request)
	ApiUsernameCheck(w http.ResponseWriter, r *http.Request)
	ApiEmailCheck(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "AdminInvite."+name)
	defer span.End()
}

func (s *svc) ApiVerifyCheck(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "ApiVerifyCheck") }
func (s *svc) ApiUsernameCheck(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "ApiUsernameCheck")
}
func (s *svc) ApiEmailCheck(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "ApiEmailCheck") }
