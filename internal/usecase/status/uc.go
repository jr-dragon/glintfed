package status

import (
	"context"

	"glintfed.org/ent/status"
	"glintfed.org/internal/data"
)

type Usecase struct {
	client *data.Client
}

func NewUsecase(client *data.Client) *Usecase {
	return &Usecase{
		client: client,
	}
}

// ObjectUrlExists
//
//	SELECT exists(*)
//	FROM statuses
//	WHERE object_url = ?
func (uc *Usecase) ObjectUrlExists(ctx context.Context, url string) (bool, error) {
	return uc.client.Ent.Status.Query().Where(status.ObjectURL(url)).Exist(ctx)
}
