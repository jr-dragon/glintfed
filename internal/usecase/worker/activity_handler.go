package worker

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"glintfed.org/ent"
	"glintfed.org/ent/notification"
	"glintfed.org/ent/profile"
	"glintfed.org/ent/status"
	"glintfed.org/ent/story"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/errs"
)

type ActivityHandler struct {
	client *data.Client

	pr ProfileRemover
}

func NewActivityHandler(client *data.Client) *ActivityHandler {
	return &ActivityHandler{
		client: client,
		pr:     NewDeletePipeline(client),
	}
}

func (ah *ActivityHandler) Dispatch(ctx context.Context, header http.Header, payload InboxPayload) error {
	if payload.Type == nil {
		return fmt.Errorf("missing payload type")
	}

	switch *payload.Type {
	case "Add":
		return ah.Add(ctx, header, payload)
	case "Create":
		return ah.Create(ctx, header, payload)
	case "Announce":
		return ah.Announce(ctx, header, payload)
	case "Accept":
		return ah.Accept(ctx, header, payload)
	case "Delete":
		return ah.Delete(ctx, header, payload)
	case "Like":
		return ah.Like(ctx, header, payload)
	case "Reject":
		return ah.Reject(ctx, header, payload)
	case "Undo":
		return ah.Undo(ctx, header, payload)
	case "View":
		return ah.View(ctx, header, payload)
	case "Story:Reaction":
		return ah.StoryReaction(ctx, header, payload)
	case "Story:Reply":
		return ah.StoryReply(ctx, header, payload)
	case "Flag":
		return ah.Flag(ctx, header, payload)
	case "Update":
		return ah.Update(ctx, header, payload)
	case "Move":
		return ah.Move(ctx, header, payload)
	default:
		return fmt.Errorf("unknown type: %s", *payload.Type)
	}
}

func (ah *ActivityHandler) Add(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Create(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Announce(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Accept(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Delete(ctx context.Context, header http.Header, payload InboxPayload) error {
	if payload.Actor == nil || payload.Object == nil {
		return nil
	}

	actor := *payload.Actor
	objID := payload.Object.ID
	objType := payload.Object.Type

	// Validate actor and object ID are valid URLs and from the same host
	actorURL, err := url.Parse(actor)
	if err != nil {
		return err
	}
	objectURL, err := url.Parse(objID)
	if err != nil {
		return err
	}

	if actorURL.Host != objectURL.Host {
		slog.WarnContext(ctx, "actor host mismatch with object host", slog.String("actor", actor), slog.String("object_id", objID))
		return nil
	}

	switch objType {
	case "Person":
		p, err := ah.client.Ent.Profile.Query().Where(profile.RemoteURL(actor)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
			return nil
		}

		if p.PrivateKey != "" {
			return nil
		}

		// Delete notifications where this actor is the actor
		if _, err := ah.client.Ent.Notification.Delete().Where(notification.ActorID(p.ID)).Exec(ctx); err != nil {
			return err
		}

		return ah.pr.RemoteProfile(ctx, p)

	case "Tombstone":
		p, err := ah.client.Ent.Profile.Query().Where(profile.RemoteURL(actor)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
			return nil
		}

		if p.PrivateKey != "" {
			return nil
		}

		s, err := ah.client.Ent.Status.Query().
			Where(status.Or(status.ObjectURL(objID), status.URL(objID))).
			First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
			return nil
		}

		if s.ProfileID != p.ID {
			return nil
		}

		// Delete notifications related to this status
		if _, err := ah.client.Ent.Notification.Delete().
			Where(
				notification.ActorID(s.ProfileID),
				notification.ItemID(s.ID),
				notification.ItemType("App\\Status"),
			).Exec(ctx); err != nil {
			return err
		}

		if (s.Scope == "public" || s.Scope == "unlisted" || s.Scope == "private") &&
			(s.Type != "story:reaction" && s.Type != "story:reply" && s.Type != "reply") {
			// TODO: FeedRemoveRemotePipeline::dispatch($status->id, $status->profile_id)->onQueue('feed');
		}

		// TODO: RemoteStatusDelete::dispatch($status)->onQueue('high');
		return nil

	case "Story":
		st, err := ah.client.Ent.Story.Query().Where(story.ObjectID(objID)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
			return nil
		}

		_ = st
		// TODO: StoryExpire::dispatch($story)->onQueue('story');
		return nil
	default:
		return nil
	}
}

func (ah *ActivityHandler) Like(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Reject(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Undo(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) View(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) StoryReaction(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) StoryReply(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Flag(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Update(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}

func (ah *ActivityHandler) Move(ctx context.Context, header http.Header, payload InboxPayload) error {
	return errs.Todo
}
