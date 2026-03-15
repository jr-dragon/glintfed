package federation

import (
	"context"
	"net/http"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	"glintfed.org/internal/service/internal"
)

type Service interface {
	SharedInbox(w http.ResponseWriter, r *http.Request)
	UserInbox(w http.ResponseWriter, r *http.Request)
	Webfinger(w http.ResponseWriter, r *http.Request)
	NodeinfoWellKnown(w http.ResponseWriter, r *http.Request)
	HostMeta(w http.ResponseWriter, r *http.Request)
	Nodeinfo(w http.ResponseWriter, r *http.Request)
}

//go:generate moq -rm -out mock_instance_usecase.go . InstanceUsecase
type InstanceUsecase interface {
	GetLocalPostsCount(ctx context.Context) (int, error)
	GetTotalUsers(ctx context.Context) (int, error)
	GetMonthActiveUsers(ctx context.Context) (int, error)
	GetHalfYearActiveUsers(ctx context.Context) (int, error)
	GetBlockedDomains(ctx context.Context) (map[string]struct{}, error)
}

//go:generate moq -rm -out mock_profile_usecase.go . ProfileUsecase
type ProfileUsecase interface {
	GetByUsername(ctx context.Context, username string) (*ent.Profile, error)
	RemoteUrlExists(ctx context.Context, url string) (bool, error)

	Url(profile *ent.Profile, surfixes ...string) string
	Permalink(profile *ent.Profile, surfixes ...string) string
}

//go:generate moq -rm -out mock_status_usecase.go . StatusUsecase
type StatusUsecase interface {
	ObjectUrlExists(ctx context.Context, url string) (bool, error)
}

func New(cfg data.Config, iuc InstanceUsecase, puc ProfileUsecase, suc StatusUsecase) Service {
	return &svc{
		cfg: cfg,

		iuc: iuc,
		puc: puc,
		suc: suc,
	}
}

type svc struct {
	cfg data.Config

	iuc InstanceUsecase
	puc ProfileUsecase
	suc StatusUsecase
}

func (s *svc) UserInbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.UserInbox")
	defer span.End()
	// TODO: Implement
}

func (s *svc) HostMeta(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.HostMeta")
	defer span.End()
	// TODO: Implement
}
