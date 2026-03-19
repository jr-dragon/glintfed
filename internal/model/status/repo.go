package status

import (
	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

type Repo struct {
	*ent.StatusClient
}

func NewRepo(client *data.Client) *Repo {
	return &Repo{
		StatusClient: client.Ent.Status,
	}
}
