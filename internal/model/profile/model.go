package profile

import (
	"glintfed.org/ent"
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
