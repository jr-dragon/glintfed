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

//go:generate go tool moq -rm -out mock_instance_model.go . InstanceModel
type InstanceModel interface {
	GetBlockedDomains(ctx context.Context) (map[string]struct{}, error)
}

//go:generate go tool moq -rm -out mock_user_model.go . UserModel
type UserModel interface {
	GetTotalUsers(ctx context.Context) (int, error)
	GetMonthActiveUsers(ctx context.Context) (int, error)
	GetHalfYearActiveUsers(ctx context.Context) (int, error)
}

//go:generate go tool moq -rm -out mock_profile_model.go . ProfileModel
type ProfileModel interface {
	GetByUsername(ctx context.Context, username string) (*ent.Profile, error)
	RemoteUrlExists(ctx context.Context, url string) (bool, error)
}

//go:generate go tool moq -rm -out mock_status_model.go . StatusModel
type StatusModel interface {
	GetLocalPostsCount(ctx context.Context) (int, error)
	ObjectUrlExists(ctx context.Context, url string) (bool, error)
}

//go:generate go tool moq -rm -out mock_worker_usecase.go . WorkerUsecase
type WorkerUsecase interface {
	Delete(ctx context.Context, header http.Header, payload worker.InboxPayload) error
	Inbox(ctx context.Context, header http.Header, payload worker.InboxPayload) error
	Validate(ctx context.Context, username string, header http.Header, payload worker.InboxPayload) error
}

func New(cfg *data.Config, im InstanceModel, pm ProfileModel, sm StatusModel, um UserModel, wuc WorkerUsecase) Service {
	return &svc{
		cfg: cfg,

		wuc: wuc,

		im: im,
		pm: pm,
		sm: sm,
		um: um,
	}
}

type svc struct {
	cfg *data.Config

	im InstanceModel
	pm ProfileModel
	sm StatusModel
	um UserModel

	wuc WorkerUsecase
}
