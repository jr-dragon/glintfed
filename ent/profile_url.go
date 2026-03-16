package ent

import (
	"strings"
)

func (p *Profile) Url(url string, suffixes ...string) string {
	if p.RemoteURL != "" {
		return p.RemoteURL
	}

	return url + "/" + p.Username + strings.Join(suffixes, "")
}

func (p *Profile) Permalink(url string, suffixes ...string) string {
	if p.RemoteURL != "" {
		return p.RemoteURL
	}

	return url + "/users/" + p.Username + strings.Join(suffixes, "")
}
