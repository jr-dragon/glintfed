package admin

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

type Service interface {
	GetAdminTabs(w http.ResponseWriter, r *http.Request)
	GetInteractionLogs(w http.ResponseWriter, r *http.Request)
	GetBlocks(w http.ResponseWriter, r *http.Request)
	ExportBlocks(w http.ResponseWriter, r *http.Request)
	AddBlock(w http.ResponseWriter, r *http.Request)
	UndoBlock(w http.ResponseWriter, r *http.Request)
	GetReportList(w http.ResponseWriter, r *http.Request)
}

func New() Service {
	return &svc{}
}

type svc struct{}

func (s *svc) GetAdminTabs(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.GetAdminTabs")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetInteractionLogs(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.GetInteractionLogs")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetBlocks(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.GetBlocks")
	defer span.End()
	// TODO: Implement
}

func (s *svc) ExportBlocks(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.ExportBlocks")
	defer span.End()
	// TODO: Implement
}

func (s *svc) AddBlock(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.AddBlock")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UndoBlock(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.UndoBlock")
	defer span.End()
	// TODO: Implement
}

func (s *svc) GetReportList(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.GetReportList")
	defer span.End()
	// TODO: Implement
}
