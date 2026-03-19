package profile

import (
	"context"

	"glintfed.org/ent"
	"glintfed.org/ent/profile"
	"glintfed.org/internal/data"
)

type Repo struct {
	*ent.ProfileClient
}

func NewRepo(client *data.Client) *Repo {
	return &Repo{
		ProfileClient: client.Ent.Profile,
	}
}

// GetByUsername
//
//	SELECT *
//	FROM profiles
//	WHERE username = ?
//	LIMIT 1
func (r *Repo) GetByUsername(ctx context.Context, username string) (*ent.Profile, error) {
	return r.Query().
		Where(profile.Username(username)).
		Only(ctx)
}

// RemoteUrlExists
//
//	SELECT exists(*)
//	FROM profiles
//	WHERE remote_url = ?
func (r *Repo) RemoteUrlExists(ctx context.Context, url string) (bool, error) {
	return r.Query().
		Where(profile.RemoteURL(url)).
		Exist(ctx)
}
