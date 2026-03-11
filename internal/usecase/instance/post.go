package instance

import (
	"context"

	"glintfed.org/ent/status"
)

// GetLocalPostCount
//
//	SELECT count(*)
//	FROM `statuses`
//	WHERE
//	  `deleted_at` IS NULL AND
//	  `local` = true
//	  `type` = "share"
func (uc *Usecase) GetLocalPostsCount(ctx context.Context) (int, error) {
	return uc.client.Ent.Status.Query().
		Where(
			status.DeletedAtIsNil(),
			status.LocalEQ(true),
			status.TypeNEQ("share"),
		).
		Count(ctx)
}
