package profile

import (
	"context"

	"glintfed.org/ent"
	"glintfed.org/ent/profile"
)

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

// GetActiveRemoteProfile
//
//	SELECT *
//	FROM profiles
//	WHERE
//		domain IS NOT NULL AND
//		status IS NULL AND
//		remote_url = ?
//	LIMIT 1
func (m *Model) GetActiveRemoteProfile(ctx context.Context, url string) (*ent.Profile, error) {
	return m.Query().
		Where(
			profile.DomainNotNil(),
			profile.StatusIsNil(),
			profile.RemoteURL(url),
		).
		First(ctx)
}

