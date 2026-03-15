package instance

import (
	"context"

	"github.com/samber/lo"
	"glintfed.org/ent"
	"glintfed.org/ent/instance"
)

func (uc *Usecase) GetBlockedDomains(ctx context.Context) (map[string]struct{}, error) {
	domains, err := uc.client.Ent.Instance.Query().
		Select("domain").
		Where(instance.Banned(true)).All(ctx)
	if err != nil {
		return nil, err
	}

	return lo.Keyify(lo.Map(domains, func(d *ent.Instance, _ int) string { return d.Domain })), nil
}
