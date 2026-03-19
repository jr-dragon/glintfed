package instance

import (
	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

type Repo struct {
	*ent.InstanceClient
}

func NewRepo(client *data.Client) *Repo {
	return &Repo{
		InstanceClient: client.Ent.Instance,
	}
}
