package status

import (
	"context"

	"glintfed.org/ent/status"
)

// ObjectUrlExists
//
//	SELECT exists(*)
//	FROM statuses
//	WHERE object_url = ?
func (r *Repo) ObjectUrlExists(ctx context.Context, url string) (bool, error) {
	return r.Query().Where(status.ObjectURL(url)).Exist(ctx)
}
