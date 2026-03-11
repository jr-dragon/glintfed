package instance

import "glintfed.org/internal/data"

type Usecase struct {
	client *data.Client
}

func NewUsecase(client *data.Client) *Usecase {
	return &Usecase{
		client: client,
	}
}
