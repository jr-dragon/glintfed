package profile

import (
	"strings"

	"glintfed.org/ent"
)

func (uc *Usecase) Url(profile *ent.Profile, surfixes ...string) string {
	if profile.RemoteURL != "" {
		return profile.RemoteURL
	}

	return profile.Username + strings.Join(surfixes, "")
}

func (uc *Usecase) Permalink(profile *ent.Profile, surfixes ...string) string {
	if profile.RemoteURL != "" {
		return profile.RemoteURL
	}

	return "users/" + profile.Username + strings.Join(surfixes, "")
}
