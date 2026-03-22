package ent

import (
	"log/slog"
	"net/url"
	"strings"
)

func (p *Profile) Url(baseUrl string, suffixes ...string) string {
	if p.RemoteURL != "" {
		return p.RemoteURL
	}

	res, err := url.JoinPath(baseUrl, p.Username, strings.Join(suffixes, ""))
	if err != nil {
		slog.Error("failed to join path",
			slog.String("baseUrl", baseUrl),
			slog.String("profile_username", p.Username),
			slog.Any("suffixes", suffixes),
		)
	}

	return res
}

func (p *Profile) Permalink(baseUrl string, suffixes ...string) string {
	if p.RemoteURL != "" {
		return p.RemoteURL
	}

	res, err := url.JoinPath(baseUrl, p.Username, strings.Join(suffixes, ""))
	if err != nil {
		slog.Error("failed to join path",
			slog.String("baseUrl", baseUrl),
			slog.String("profile_username", p.Username),
			slog.Any("suffixes", suffixes),
		)
	}

	return url + "/users/" + p.Username + strings.Join(suffixes, "")
}
