package worker

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"

	"glintfed.org/ent"
	"glintfed.org/ent/notification"
	"glintfed.org/ent/profile"
	"glintfed.org/ent/status"
	"glintfed.org/ent/story"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
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

func (ah *ActivityHandler) Dispatch(ctx context.Context, header http.Header, payload InboxPayload) {
	
}

func (ah *ActivityHandler) Delete(ctx context.Context, header http.Header, payload InboxPayload) {
	if payload.Actor == nil || payload.Object == nil {
		return
	}

	actor := *payload.Actor
	objID := payload.Object.ID
	objType := payload.Object.Type

	// Validate actor and object ID are valid URLs and from the same host
	actorURL, err := url.Parse(actor)
	if err != nil {
		return
	}
	objectURL, err := url.Parse(objID)
	if err != nil {
		return
	}

	if actorURL.Host != objectURL.Host {
		slog.WarnContext(ctx, "actor host mismatch with object host", slog.String("actor", actor), slog.String("object_id", objID))
		return
	}

	switch objType {
	case "Person":
		p, err := ah.client.Ent.Profile.Query().Where(profile.RemoteURL(actor)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				slog.ErrorContext(ctx, "failed to get profile for delete", logs.ErrAttr(err), slog.String("actor", actor))
			}
			return
		}

		if p.PrivateKey != "" {
			return
		}

		// Delete notifications where this actor is the actor
		if _, err := ah.client.Ent.Notification.Delete().Where(notification.ActorID(p.ID)).Exec(ctx); err != nil {
			slog.ErrorContext(ctx, "failed to delete notifications for actor", logs.ErrAttr(err), slog.Uint64("profile_id", p.ID))
		}

		ah.pr.RemoteProfile(ctx, p)

	case "Tombstone":
		p, err := ah.client.Ent.Profile.Query().Where(profile.RemoteURL(actor)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				slog.ErrorContext(ctx, "failed to get profile for tombstone", logs.ErrAttr(err), slog.String("actor", actor))
			}
			return
		}

		if p.PrivateKey != "" {
			return
		}

		s, err := ah.client.Ent.Status.Query().
			Where(status.Or(status.ObjectURL(objID), status.URL(objID))).
			First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				slog.ErrorContext(ctx, "failed to get status for tombstone", logs.ErrAttr(err), slog.String("object_id", objID))
			}
			return
		}

		if s.ProfileID != p.ID {
			return
		}

		// Delete notifications related to this status
		if _, err := ah.client.Ent.Notification.Delete().
			Where(
				notification.ActorID(s.ProfileID),
				notification.ItemID(s.ID),
				notification.ItemType("App\\Status"),
			).Exec(ctx); err != nil {
			slog.ErrorContext(ctx, "failed to delete notifications for status tombstone", logs.ErrAttr(err), slog.Uint64("status_id", s.ID))
		}

		if (s.Scope == "public" || s.Scope == "unlisted" || s.Scope == "private") &&
			(s.Type != "story:reaction" && s.Type != "story:reply" && s.Type != "reply") {
			// TODO: FeedRemoveRemotePipeline::dispatch($status->id, $status->profile_id)->onQueue('feed');
		}

		// TODO: RemoteStatusDelete::dispatch($status)->onQueue('high');

	case "Story":
		st, err := ah.client.Ent.Story.Query().Where(story.ObjectID(objID)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				slog.ErrorContext(ctx, "failed to get story for delete", logs.ErrAttr(err), slog.String("object_id", objID))
			}
			return
		}

		_ = st
		// TODO: StoryExpire::dispatch($story)->onQueue('story');
	default:
		return
	}
}
