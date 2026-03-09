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

func (s *svc) Report(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.Report")
	defer span.End()
	// TODO: Implement
}

func (s *svc) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.DeleteAvatar")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountPosts(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.AccountPosts")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountChangePassword(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.AccountChangePassword")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountLoginActivity(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.AccountLogicActivity")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountTwoFactor(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.AccountTwoFactor")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountEmailsFromPixelfed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.AccountEmailsFromPixelfed")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountApps(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.AccountApps")
	defer span.End()
	// TODO: Implement
}

func (s *svc) InAppRegistrationPreFlightCheck(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.InAppRegistrationPreFlightCheck")
	defer span.End()
	// TODO: Implement
}

func (s *svc) InAppRegistration(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.InAppRegistration")
	defer span.End()
	// TODO: Implement
}

func (s *svc) InAppRegistrationEmailRedirect(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.InAppRegistrationEmailRedirect")
	defer span.End()
	// TODO: Implement
}

func (s *svc) InAppRegistrationConfirm(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.InAppRegistrationConfirm")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Archive(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.Archive")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Unarchive(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.Unarchive")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ArchivedPosts(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.ArchivedPosts")
	defer span.End()
	// TODO: Implement
}

func (s *svc) PlacesById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.PlacesById")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ModeratePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.ModeratePost")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetWebSettings(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.GetWebSettings")
	defer span.End()
	// TODO: Implement
}

func (s *svc) SetWebSettings(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.SetWebSettings")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetMutualAccounts(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.GetMutalAccounts")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AccountUsernameToId(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.AccountUsernameToId")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetPushState(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.GetPushState")
	defer span.End()
	// TODO: Implement
}

func (s *svc) DisablePush(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.DisablePush")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ComparePush(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.ComparePush")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UpdatePush(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.UpdatePush")
	defer span.End()
	// TODO: Implement
}

func (s *svc) StatusCreate(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.StatusCreate")
	defer span.End()
	// TODO: Implement
}

func (s *svc) NagState(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1_1.NagState")
	defer span.End()
	// TODO: Implement
}
