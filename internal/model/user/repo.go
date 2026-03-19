package user

import (
	"context"
	"time"

	"glintfed.org/ent"
	"glintfed.org/ent/user"
	"glintfed.org/internal/data"
)

type Repo struct {
	*ent.UserClient
}

func NewRepo(client *data.Client) *Repo {
	return &Repo{
		UserClient: client.Ent.User,
	}
}

// GetTotalUsers
//
//	SELECT count(*)
//	FROM users
func (r *Repo) GetTotalUsers(ctx context.Context) (int, error) {
	return r.Query().Count(ctx)
}

// GetMonthActiveUsers
//
//	SELECT count(`last_active_at`, `updated_at`)
//	FROM `users`
//	WHERE
//	  `updated_at` > datetime(NOW(), '-5 weeks') OR `last_active_at` > datetime(NOW(), '-5 weeks')
func (r *Repo) GetMonthActiveUsers(ctx context.Context) (int, error) {
	threshold := time.Now().Add(-5 * 7 * 24 * time.Hour)
	return r.Query().
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
func (r *Repo) GetHalfYearActiveUsers(ctx context.Context) (int, error) {
	threshold := time.Now().AddDate(0, -6, 0)
	return r.Query().
		Where(
			user.Or(
				user.UpdatedAtGT(threshold),
				user.LastActiveAtGT(threshold),
			),
		).
		Count(ctx)
}
