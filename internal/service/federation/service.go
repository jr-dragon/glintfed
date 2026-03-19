package federation

import (
	"context"
	"net/http"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	"glintfed.org/internal/usecase/worker"
)

type Service interface {
	SharedInbox(w http.ResponseWriter, r *http.Request)
	UserInbox(w http.ResponseWriter, r *http.Request)
	Webfinger(w http.ResponseWriter, r *http.Request)
	NodeinfoWellKnown(w http.ResponseWriter, r *http.Request)
	HostMeta(w http.ResponseWriter, r *http.Request)
	Nodeinfo(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_instance_usecase.go . InstanceUsecase
type InstanceUsecase interface {
	GetLocalPostsCount(ctx context.Context) (int, error)
	GetTotalUsers(ctx context.Context) (int, error)
	GetMonthActiveUsers(ctx context.Context) (int, error)
	GetHalfYearActiveUsers(ctx context.Context) (int, error)
	GetBlockedDomains(ctx context.Context) (map[string]struct{}, error)
}

//go:generate go tool moq -rm -out mock_profile_usecase.go . ProfileUsecase
type ProfileUsecase interface {
	GetByUsername(ctx context.Context, username string) (*ent.Profile, error)
	RemoteUrlExists(ctx context.Context, url string) (bool, error)
}

//go:generate go tool moq -rm -out mock_status_usecase.go . StatusUsecase
type StatusUsecase interface {
	ObjectUrlExists(ctx context.Context, url string) (bool, error)
}

//go:generate go tool moq -rm -out mock_worker_usecase.go . WorkerUsecase
type WorkerUsecase interface {
	Delete(ctx context.Context, header http.Header, payload worker.InboxPayload) error
	Inbox(ctx context.Context, header http.Header, payload worker.InboxPayload) error
	Validate(ctx context.Context, username string, header http.Header, payload worker.InboxPayload) error
}

func New(cfg *data.Config, iuc InstanceUsecase, puc ProfileUsecase, suc StatusUsecase, wuc WorkerUsecase) Service {
	return &svc{
		cfg: cfg,

		iuc: iuc,
		puc: puc,
		suc: suc,
		wuc: wuc,
	}
}

type svc struct {
	cfg *data.Config

	iuc InstanceUsecase
	puc ProfileUsecase
	suc StatusUsecase
	wuc WorkerUsecase
}
