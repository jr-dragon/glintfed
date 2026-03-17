package worker

import (
	"context"
	"log/slog"

	"glintfed.org/ent"
	"glintfed.org/ent/avatar"
	"glintfed.org/ent/conversation"
	"glintfed.org/ent/directmessage"
	"glintfed.org/ent/follower"
	"glintfed.org/ent/followrequest"
	"glintfed.org/ent/like"
	"glintfed.org/ent/mediatag"
	"glintfed.org/ent/mention"
	"glintfed.org/ent/notification"
	"glintfed.org/ent/poll"
	"glintfed.org/ent/pollvote"
	"glintfed.org/ent/report"
	"glintfed.org/ent/story"
	"glintfed.org/ent/storyview"
	"glintfed.org/ent/userfilter"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
)

type DeletePipeline struct {
	client *data.Client
}

func NewDeletePipeline(client *data.Client) *DeletePipeline {
	return &DeletePipeline{
		client: client,
	}
}

func (dp *DeletePipeline) RemoteProfile(ctx context.Context, profile *ent.Profile) {
	if profile == nil {
		return
	}

	// Local profile or profile with private key should not be processed by this worker
	if profile.Domain == "" || profile.PrivateKey != "" {
		return
	}

	pid := profile.ID

	// TODO:
	// AccountService::del($pid);

	// TODO:
	// Status::whereProfileId($pid)
	// ->chunk(50, function ($statuses) {
	//     foreach ($statuses as $status) {
	//         RemoteStatusDelete::dispatch($status)->onQueue('delete');
	//     }
	// });

	// 2. Delete Polls & Poll Votes
	if _, err := dp.client.Ent.PollVote.Delete().Where(pollvote.ProfileID(pid)).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete poll vote", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}
	if _, err := dp.client.Ent.Poll.Delete().Where(poll.ProfileID(pid)).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete poll", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 3. Delete Avatar
	if _, err := dp.client.Ent.Avatar.Delete().Where(avatar.ProfileID(pid)).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete avatar", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 4. Delete Media Tags
	if _, err := dp.client.Ent.MediaTag.Delete().Where(mediatag.ProfileID(pid)).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete media tag", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 5. Delete DMs & Conversations
	if _, err := dp.client.Ent.DirectMessage.Delete().Where(
		directmessage.Or(
			directmessage.FromID(pid),
			directmessage.ToID(pid),
		),
	).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete direct messages", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}
	if _, err := dp.client.Ent.Conversation.Delete().Where(
		conversation.Or(
			conversation.FromID(pid),
			conversation.ToID(pid),
		),
	).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete conversations", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 6. Delete Follow Requests
	if _, err := dp.client.Ent.FollowRequest.Delete().Where(
		followrequest.Or(
			followrequest.FollowingID(pid),
			followrequest.FollowerID(pid),
		),
	).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete follow requests", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 7. Delete Followers
	if _, err := dp.client.Ent.Follower.Delete().Where(
		follower.Or(
			follower.ProfileID(pid),
			follower.FollowingID(pid),
		),
	).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete followers", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 8. Delete Likes
	if _, err := dp.client.Ent.Like.Delete().Where(like.ProfileID(pid)).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete likes", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 9. Delete Stories & Story Views
	if _, err := dp.client.Ent.StoryView.Delete().Where(storyview.ProfileID(pid)).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete story views", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	if _, err := dp.client.Ent.Story.Delete().Where(story.ProfileID(pid)).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete stories", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}
	// TODO:
	//  foreach ($stories as $story) {
	//         $path = storage_path('app/'.$story->path);
	//         if (is_file($path)) {
	//             unlink($path);
	//         }
	//         $story->forceDelete();
	//     }

	// 10. Delete User Filters (mutes/blocks)
	if _, err := dp.client.Ent.UserFilter.Delete().Where(
		userfilter.And(
			userfilter.FilterableType("App\\Profile"),
			userfilter.FilterableID(pid),
		),
	).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete user filters", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 11. Delete Mentions
	if _, err := dp.client.Ent.Mention.Delete().Where(mention.ProfileID(pid)).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete mentions", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 12. Delete Notifications
	if _, err := dp.client.Ent.Notification.Delete().Where(
		notification.Or(
			notification.ProfileID(pid),
			notification.ActorID(pid),
		),
	).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete notifications", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 13. Delete Reports
	if _, err := dp.client.Ent.Report.Delete().Where(
		report.Or(
			report.ProfileID(pid),
			report.ReportedProfileID(pid),
		),
	).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete reports", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}

	// 14. Finalize: Delete Profile
	if err := dp.client.Ent.Profile.DeleteOne(profile).Exec(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to delete profile", logs.ErrAttr(err), slog.Uint64("profile_id", pid))
	}
}
