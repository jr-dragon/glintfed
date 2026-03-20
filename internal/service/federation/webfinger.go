package federation

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

type SharedInboxResponse struct {
	Subject string            `json:"subject"`
	Aliases []string          `json:"aliases"`
	Links   []SharedInboxLink `json:"links"`
}

type SharedInboxLink struct {
	Rel      string  `json:"rel"`
	Type     *string `json:"type,omitzero"`
	Href     *string `json:"href,omitzero"`
	Template *string `json:"template,omitzero"`
}

func (s *svc) Webfinger(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.Webfinger")
	defer span.End()

	resource := r.URL.Query().Get("resource")
	if !s.cfg.App.Federation.Webfinger.Enabled || resource == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	u, err := url.Parse(s.cfg.App.Url)
	if err != nil || u == nil {
		const msg = "failed to parse app url"
		slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if isShareboxResource(s.cfg, u, resource) {
		resp := SharedInboxResponse{
			Subject: resource,
			Aliases: []string{"https://" + u.Host + "/i/actor"},
			Links: []SharedInboxLink{
				{
					Rel:  "http://webfinger.net/rel/profile-page",
					Type: new("text/html"),
					Href: new("https://" + u.Host + "/site/kb/instance-actor"),
				},
				{
					Rel:  "self",
					Type: new("application/activity+json"),
					Href: new("https://" + u.Host + "/i/actor"),
				},
				{
					Rel:      "http://ostatus.org/schema/1.0/subscribe",
					Template: new("https://" + u.Host + "/authorize_interaction?uri={uri}"),
				},
			},
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	var username string
	if after, ok := strings.CutPrefix(resource, "https://"+u.Host+"/users/"); ok {
		username = after
	} else if after, ok := strings.CutPrefix(resource, "acct:"); ok {
		parts := strings.Split(after, "@")
		if len(parts) != 2 || parts[1] != u.Host {
			slog.ErrorContext(r.Context(), "invalid acct resource", slog.String("resource", resource))
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		username = parts[0]
	}

	if username != "" {
		if len(username) > s.cfg.App.MaxNameLength {
			slog.ErrorContext(r.Context(), "username too long", slog.String("username", username))
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		if !isValidUsername(username) {
			slog.ErrorContext(r.Context(), "invalid username", slog.String("username", username))
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		profile, err := s.pm.GetByUsername(r.Context(), username)
		if err != nil {
			slog.ErrorContext(r.Context(), "failed to get profile by username", logs.ErrAttr(err))
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		webfinger, err := s.newWebfinger(profile)
		if err != nil {
			slog.ErrorContext(r.Context(), "failed to create webfinger struct", logs.ErrAttr(err))
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(webfinger)
		return
	}

	slog.ErrorContext(r.Context(), "invalid resource format", slog.String("resource", resource))
	http.Error(w, "", http.StatusBadRequest)
}

func isShareboxResource(cfg *data.Config, u *url.URL, resource string) bool {
	var sb strings.Builder
	sb.WriteString("acct:")
	sb.WriteString(u.Host)
	sb.WriteRune('@')
	sb.WriteString(u.Host)

	return cfg.App.Federation.Activitypub.SharedInbox &&
		resource == sb.String()
}

func isValidUsername(username string) bool {
	for _, c := range username {
		if c == '_' || c == '.' || c == '-' {
			continue
		}
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			continue
		}

		return false
	}

	return true
}

type webfinger struct {
	Subject string            `json:"subject"`
	Aliases []string          `json:"aliases"`
	Links   []SharedInboxLink `json:"links"`
}

func (s *svc) newWebfinger(profile *ent.Profile) (webfinger, error) {
	u, err := url.Parse(s.cfg.App.Url)
	if err != nil || u == nil {
		slog.Error("failed to parse app url", logs.ErrAttr(err), slog.Any("parsed", u))
		return webfinger{}, err
	}

	return webfinger{
		Subject: fmt.Sprintf("acct:%s@%s", profile.Username, u.Host),
		Aliases: []string{
			profile.Url(s.cfg.App.Url),
			profile.Permalink(s.cfg.App.Url),
		},
		Links: []SharedInboxLink{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: new("text/html"),
				Href: new(profile.Url(s.cfg.App.Url)),
			},
			{
				Rel:  "http://schemas.google.com/g/2010#updates-from",
				Type: new("application/atom+xml"),
				Href: new(profile.Permalink(s.cfg.App.Url, ".atom")),
			},
			{
				Rel:  "self",
				Type: new("application/activity+json"),
				Href: new(profile.Permalink(s.cfg.App.Url)),
			},
			{
				Rel:  "http://webfinger.net/rel/avatar",
				Type: new("image/webp"),
				Href: &profile.AvatarURL,
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: new("https://" + u.Host + "/authorize_interaction?uri={uri}"),
			},
		},
	}, nil
}
