package status

import (
	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

type Model struct {
	*ent.StatusClient
}

func NewModel(client *data.Client) *Model {
	return &Model{
		StatusClient: client.Ent.Status,
	}
}
