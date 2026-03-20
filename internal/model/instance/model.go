package instance

import (
	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

type Model struct {
	*ent.InstanceClient
}

func NewModel(client *data.Client) *Model {
	return &Model{
		InstanceClient: client.Ent.Instance,
	}
}
