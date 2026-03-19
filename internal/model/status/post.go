package status

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
func (m *Model) GetLocalPostsCount(ctx context.Context) (int, error) {
	return m.Query().
		Where(
			status.DeletedAtIsNil(),
			status.Local(true),
			status.TypeNEQ("share"),
		).
		Count(ctx)
}
