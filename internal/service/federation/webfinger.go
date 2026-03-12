package federation

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

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
	Type     *string `json:"type,omitempty"`
	Href     *string `json:"href,omitempty"`
	Template *string `json:"template,omitempty"`
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

	if u.Scheme == "https" {
		if strings.HasPrefix(resource, "https://"+u.Host+"/users/") {
			username := strings.TrimLeft(resource, "https://"+u.Host+"/users/")
			if len(username) > s.cfg.App.MaxNameLength {
				slog.ErrorContext(r.Context(), "username too long")
				http.Error(w, "", http.StatusBadRequest)
				return
			}

		} else {
			slog.ErrorContext(r.Context(), "resource starts with 'https://', but invalid domain or path")
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}
}

func isShareboxResource(cfg data.Config, u *url.URL, resource string) bool {
	var sb strings.Builder
	sb.WriteString("acct:")
	sb.WriteString(u.Host)
	sb.WriteRune('@')
	sb.WriteString(u.Host)

	return cfg.App.Federation.Activitypub.SharedInbox &&
		resource == sb.String()
}
