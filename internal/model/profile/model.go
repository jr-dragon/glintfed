package profile

import (
	"context"

	"glintfed.org/ent"
	"glintfed.org/ent/profile"
	"glintfed.org/internal/data"
)

type Model struct {
	*ent.ProfileClient
}

func NewModel(client *data.Client) *Model {
	return &Model{
		ProfileClient: client.Ent.Profile,
	}
}

// GetByUsername
//
//	SELECT *
//	FROM profiles
//	WHERE username = ?
//	LIMIT 1
func (m *Model) GetByUsername(ctx context.Context, username string) (*ent.Profile, error) {
	return m.Query().
		Where(profile.Username(username)).
		Only(ctx)
}

// RemoteUrlExists
//
//	SELECT exists(*)
//	FROM profiles
//	WHERE remote_url = ?
func (m *Model) RemoteUrlExists(ctx context.Context, url string) (bool, error) {
	return m.Query().
		Where(profile.RemoteURL(url)).
		Exist(ctx)
}
