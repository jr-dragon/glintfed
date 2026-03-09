package apiv1dot1

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Report(w http.ResponseWriter, r *http.Request)
	DeleteAvatar(w http.ResponseWriter, r *http.Request)
	AccountPosts(w http.ResponseWriter, r *http.Request)
	AccountChangePassword(w http.ResponseWriter, r *http.Request)
	AccountLoginActivity(w http.ResponseWriter, r *http.Request)
	AccountTwoFactor(w http.ResponseWriter, r *http.Request)
	AccountEmailsFromPixelfed(w http.ResponseWriter, r *http.Request)
	AccountApps(w http.ResponseWriter, r *http.Request)
	InAppRegistrationPreFlightCheck(w http.ResponseWriter, r *http.Request)
	InAppRegistration(w http.ResponseWriter, r *http.Request)
	InAppRegistrationEmailRedirect(w http.ResponseWriter, r *http.Request)
	InAppRegistrationConfirm(w http.ResponseWriter, r *http.Request)
	Archive(w http.ResponseWriter, r *http.Request)
	Unarchive(w http.ResponseWriter, r *http.Request)
	ArchivedPosts(w http.ResponseWriter, r *http.Request)
	PlacesById(w http.ResponseWriter, r *http.Request)
	ModeratePost(w http.ResponseWriter, r *http.Request)
	GetWebSettings(w http.ResponseWriter, r *http.Request)
	SetWebSettings(w http.ResponseWriter, r *http.Request)
	GetMutualAccounts(w http.ResponseWriter, r *http.Request)
	AccountUsernameToId(w http.ResponseWriter, r *http.Request)
	GetPushState(w http.ResponseWriter, r *http.Request)
	DisablePush(w http.ResponseWriter, r *http.Request)
	ComparePush(w http.ResponseWriter, r *http.Request)
	UpdatePush(w http.ResponseWriter, r *http.Request)
	StatusCreate(w http.ResponseWriter, r *http.Request)
	NagState(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "ApiV1Dot1."+name)
	defer span.End()
}

func (s *svc) Report(w http.ResponseWriter, r *http.Request)       { s.stub(w, r, "Report") }
func (s *svc) DeleteAvatar(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "DeleteAvatar") }
func (s *svc) AccountPosts(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "AccountPosts") }
func (s *svc) AccountChangePassword(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountChangePassword")
}
func (s *svc) AccountLoginActivity(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountLoginActivity")
}
func (s *svc) AccountTwoFactor(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountTwoFactor")
}
func (s *svc) AccountEmailsFromPixelfed(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountEmailsFromPixelfed")
}
func (s *svc) AccountApps(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "AccountApps") }
func (s *svc) InAppRegistrationPreFlightCheck(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "InAppRegistrationPreFlightCheck")
}
func (s *svc) InAppRegistration(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "InAppRegistration")
}
func (s *svc) InAppRegistrationEmailRedirect(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "InAppRegistrationEmailRedirect")
}
func (s *svc) InAppRegistrationConfirm(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "InAppRegistrationConfirm")
}
func (s *svc) Archive(w http.ResponseWriter, r *http.Request)        { s.stub(w, r, "Archive") }
func (s *svc) Unarchive(w http.ResponseWriter, r *http.Request)      { s.stub(w, r, "Unarchive") }
func (s *svc) ArchivedPosts(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "ArchivedPosts") }
func (s *svc) PlacesById(w http.ResponseWriter, r *http.Request)     { s.stub(w, r, "PlacesById") }
func (s *svc) ModeratePost(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "ModeratePost") }
func (s *svc) GetWebSettings(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetWebSettings") }
func (s *svc) SetWebSettings(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "SetWebSettings") }
func (s *svc) GetMutualAccounts(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "GetMutualAccounts")
}
func (s *svc) AccountUsernameToId(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "AccountUsernameToId")
}
func (s *svc) GetPushState(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetPushState") }
func (s *svc) DisablePush(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "DisablePush") }
func (s *svc) ComparePush(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "ComparePush") }
func (s *svc) UpdatePush(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "UpdatePush") }
func (s *svc) StatusCreate(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "StatusCreate") }
func (s *svc) NagState(w http.ResponseWriter, r *http.Request)     { s.stub(w, r, "NagState") }
