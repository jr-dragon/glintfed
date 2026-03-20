package profile

import (
	"context"

	"glintfed.org/ent"
	"glintfed.org/ent/profile"
)

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

// GetActiveLocalProfile
//
//	SELECT *
//	FROM profiles
//	WHERE
//		username = ? AND
//		domain IS NULL AND
//		status IS NULL
//	LIMIT 1
func (m *Model) GetActiveLocalProfile(ctx context.Context, username string) (*ent.Profile, error) {
	return m.Query().
		Where(
			profile.Username(username),
			profile.DomainIsNil(),
			profile.StatusIsNil(),
		).
		First(ctx)
}

// GetByKeyID
//
//	SELECT *
//	FROM profiles
//	WHERE key_id = ?
//	LIMIT 1
func (m *Model) GetByKeyID(ctx context.Context, kid string) (*ent.Profile, error) {
	return m.Query().
		Where(profile.KeyID(kid)).
		First(ctx)
}
