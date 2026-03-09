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

func (s *svc) stub(w http.ResponseWriter, r *http.Request, name string) {
	_, span := internal.T.Start(r.Context(), "AdminApi."+name)
	defer span.End()
}

func (s *svc) Supported(w http.ResponseWriter, r *http.Request)      { s.stub(w, r, "Supported") }
func (s *svc) GetStats(w http.ResponseWriter, r *http.Request)       { s.stub(w, r, "GetStats") }
func (s *svc) Autospam(w http.ResponseWriter, r *http.Request)       { s.stub(w, r, "Autospam") }
func (s *svc) AutospamHandle(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "AutospamHandle") }
func (s *svc) ModReports(w http.ResponseWriter, r *http.Request)     { s.stub(w, r, "ModReports") }
func (s *svc) ModReportHandle(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "ModReportHandle")
}
func (s *svc) GetConfiguration(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "GetConfiguration")
}
func (s *svc) UpdateConfiguration(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "UpdateConfiguration")
}
func (s *svc) GetUsers(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetUsers") }
func (s *svc) GetUser(w http.ResponseWriter, r *http.Request)  { s.stub(w, r, "GetUser") }
func (s *svc) UserAdminAction(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "UserAdminAction")
}
func (s *svc) Instances(w http.ResponseWriter, r *http.Request)   { s.stub(w, r, "Instances") }
func (s *svc) GetInstance(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetInstance") }
func (s *svc) ModerateInstance(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "ModerateInstance")
}
func (s *svc) RefreshInstanceStats(w http.ResponseWriter, r *http.Request) {
	s.stub(w, r, "RefreshInstanceStats")
}
func (s *svc) GetAllStats(w http.ResponseWriter, r *http.Request) { s.stub(w, r, "GetAllStats") }
