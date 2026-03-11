package federation

import (
	"context"
	"net/http"

	"glintfed.org/internal/data"
	"glintfed.org/internal/service/internal"
	"glintfed.org/internal/usecase/instance"
)

type Service interface {
	SharedInbox(w http.ResponseWriter, r *http.Request)
	UserInbox(w http.ResponseWriter, r *http.Request)
	Webfinger(w http.ResponseWriter, r *http.Request)
	NodeinfoWellKnown(w http.ResponseWriter, r *http.Request)
	HostMeta(w http.ResponseWriter, r *http.Request)
	Nodeinfo(w http.ResponseWriter, r *http.Request)
}

type InstanceUsecase interface {
	GetLocalPostsCount(ctx context.Context) (int, error)
	GetTotalUsers(ctx context.Context) (int, error)
	GetMonthActiveUsers(ctx context.Context) (int, error)
	GetHalfYearActiveUsers(ctx context.Context) (int, error)
}

func New(cfg data.Config, client *data.Client) Service {
	return &svc{
		cfg: cfg,

		iuc: instance.NewUsecase(client),
	}
}

type svc struct {
	cfg data.Config

	iuc InstanceUsecase
}

func (s *svc) SharedInbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.SharedInbox")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UserInbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.UserInbox")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Webfinger(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.Webfinger")
	defer span.End()
	// TODO: Implement
}

func (s *svc) HostMeta(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.HostMeta")
	defer span.End()
	// TODO: Implement
}
