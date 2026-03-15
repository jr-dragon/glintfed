package worker

import (
	"context"
	"net/http"

	"glintfed.org/internal/data"
)

type InboxUsecase struct {
	client *data.Client
}

func NewInboxUsecase(client *data.Client) *InboxUsecase {
	return &InboxUsecase{
		client: client,
	}
}

func (inbox *InboxUsecase) Delete(ctx context.Context, header http.Header, payload any) {

}

func (inbox *InboxUsecase) Inbox(ctx context.Context, header http.Header, payload any) {

}

func (inbox *InboxUsecase) Validate(ctx context.Context, username string, header http.Header, payload any) {

}
