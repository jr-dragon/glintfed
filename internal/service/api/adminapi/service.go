package adminapi

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	Supported(w http.ResponseWriter, r *http.Request)
	GetStats(w http.ResponseWriter, r *http.Request)
	Autospam(w http.ResponseWriter, r *http.Request)
	AutospamHandle(w http.ResponseWriter, r *http.Request)
	ModReports(w http.ResponseWriter, r *http.Request)
	ModReportHandle(w http.ResponseWriter, r *http.Request)
	GetConfiguration(w http.ResponseWriter, r *http.Request)
	UpdateConfiguration(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UserAdminAction(w http.ResponseWriter, r *http.Request)
	Instances(w http.ResponseWriter, r *http.Request)
	GetInstance(w http.ResponseWriter, r *http.Request)
	ModerateInstance(w http.ResponseWriter, r *http.Request)
	RefreshInstanceStats(w http.ResponseWriter, r *http.Request)
	GetAllStats(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) Supported(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.Supported")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetStats(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.GetStats")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Autospam(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.Autospam")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AutospamHandle(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.AutospamHandle")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ModReports(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.ModReports")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ModReportHandle(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.ModReportHandle")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetConfiguration(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.GetConfiguration")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UpdateConfiguration(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.UpdateConfiguration")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetUsers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.GetUsers")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetUser(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.GetUser")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UserAdminAction(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.UserAdminAction")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Instances(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.Instances")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetInstance(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.GetInstance")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ModerateInstance(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.ModerateInstance")
	defer span.End()
	// TODO: Implement
}

func (s *svc) RefreshInstanceStats(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.RefreshInstanceStats")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetAllStats(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AdminInvite.GetAllStats")
	defer span.End()
	// TODO: Implement
}
