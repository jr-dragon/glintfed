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
func (m *Model) ObjectUrlExists(ctx context.Context, url string) (bool, error) {
	return m.Query().Where(status.ObjectURL(url)).Exist(ctx)
}
