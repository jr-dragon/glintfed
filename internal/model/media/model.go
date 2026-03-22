package media

import (
	"context"

	"glintfed.org/ent"
	"glintfed.org/ent/media"
	"glintfed.org/internal/data"
)

type Model struct {
	*ent.MediaClient
}

func NewModel(client *data.Client) *Model {
	return &Model{
		MediaClient: client.Ent.Media,
	}
}

func (m *Model) GetCDNUrl(ctx context.Context, path string) (string, error) {
	res, err := m.Query().
		Select("cdn_url").
		Where(
			media.MediaPath(path),
			media.CdnURLNotNil(),
		).First(ctx)
	if err != nil {
		return "", err
	}

	return res.CdnURL, nil
}
