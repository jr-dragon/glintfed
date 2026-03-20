package instanceactor

import (
	"context"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

type Model struct {
	*ent.InstanceActorClient
}

func NewModel(client *data.Client) *Model {
	return &Model{
		InstanceActorClient: client.Ent.InstanceActor,
	}
}

// Get
//
//	SELECT *
//	FROM instance_actors
//	LIMIT 1
func (m *Model) Get(ctx context.Context) (*ent.InstanceActor, error) {
	return m.Query().First(ctx)
}
