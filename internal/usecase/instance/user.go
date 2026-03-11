package instance

import (
	"context"
	"time"

	"glintfed.org/ent/user"
)

// GetTotalUsers
//
//	SELECT count(*)
//	FROM users
func (s *Usecase) GetTotalUsers(ctx context.Context) (int, error) {
	return s.client.Ent.User.Query().Count(ctx)
}

// GetMonthActiveUsers
//
//	SELECT count(`last_active_at`, `updated_at`)
//	FROM `users`
//	WHERE
//	  `updated_at` > datetime(NOW(), '-5 weeks') OR `last_active_at` > datetime(NOW(), '-5 weeks')
func (s *Usecase) GetMonthActiveUsers(ctx context.Context) (int, error) {
	threshold := time.Now().Add(-5 * 7 * 24 * time.Hour)
	return s.client.Ent.User.Query().
		Where(
			user.Or(
				user.UpdatedAtGT(threshold),
				user.LastActiveAtGT(threshold),
			),
		).
		Count(ctx)
}

// GetHalfYearActiveUsers
//
//	SELECT count(`last_active_at`, `updated_at`)
//	FROM `users`
//	WHERE
//		`updated_at` > datetime(NOW(), '-6 months') OR `last_active_at` > datetime(NOW(), '-6 months')
func (s *Usecase) GetHalfYearActiveUsers(ctx context.Context) (int, error) {
	threshold := time.Now().AddDate(0, -6, 0)
	return s.client.Ent.User.Query().
		Where(
			user.Or(
				user.UpdatedAtGT(threshold),
				user.LastActiveAtGT(threshold),
			),
		).
		Count(ctx)
}
