package profile

import (
	"context"

	"glintfed.org/ent"
	"glintfed.org/ent/profile"
	"glintfed.org/internal/data"
)

type Usecase struct {
	client *data.Client
	cfg    *data.Config
}

func NewUsecase(client *data.Client, cfg *data.Config) *Usecase {
	return &Usecase{
		client: client,
		cfg:    cfg,
	}
}

// GetByUsername
//
//	SELECT *
//	FROM profiles
//	WHERE username = ?
//	LIMIT 1
func (uc *Usecase) GetByUsername(ctx context.Context, username string) (*ent.Profile, error) {
	return uc.client.Ent.Profile.Query().
		Where(profile.Username(username)).
		Only(ctx)
}

// RemoteUrlExists
//
//	SELECT exists(*)
//	FROM profiles
//	WHERE remote_url = ?
func (uc *Usecase) RemoteUrlExists(ctx context.Context, url string) (bool, error) {
	return uc.client.Ent.Profile.Query().
		Where(profile.RemoteURL(url)).
		Exist(ctx)
}
