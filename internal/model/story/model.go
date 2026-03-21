package story

import (
	"context"

	"glintfed.org/ent"
	"glintfed.org/ent/profile"
	"glintfed.org/ent/story"
	"glintfed.org/internal/data"
)

type Model struct {
	client *data.Client
}

func NewModel(client *data.Client) *Model {
	return &Model{
		client: client,
	}
}

// GetByUsernameAndID
//
//	SELECT *
//	FROM stories
//	INNER JOIN profiles ON stories.profile_id = profiles.id
//	WHERE profiles.username = ?
//		AND profiles.domain IS NULL
//		AND stories.id = ?
//		AND stories.active = true
//	LIMIT 1
func (m *Model) GetByUsernameAndID(ctx context.Context, username string, id uint64) (*ent.Story, error) {
	return m.client.Ent.Profile.Query().
		Where(
			profile.Username(username),
			profile.DomainIsNil(),
		).
		QueryStories().
		Where(
			story.ID(id),
			story.Active(true),
		).
		WithProfile().
		First(ctx)
}
