package appregister

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	VerifyCode(w http.ResponseWriter, r *http.Request)
	Onboarding(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) VerifyCode(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AppRegister.VerifyCode")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Onboarding(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AppRegister.Onboarding")
	defer span.End()
	// TODO: Implement
}
