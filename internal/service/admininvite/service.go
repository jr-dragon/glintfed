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

}

func (s *svc) ApiVerifyCheck(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AdminInvite.ApiVerifyCheck")
	defer span.End()
	// TODO: Implement
}
func (s *svc) ApiUsernameCheck(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AdminInvite.ApiUsernameCheck")
	defer span.End()
	// TODO: Implement
}
func (s *svc) ApiEmailCheck(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AdminInvite.ApiEmailCheck")
	defer span.End()
	// TODO: Implement
}
