package federation

import (
	"encoding/json"
	"net/http"
	"time"

	"glintfed.org/ent/status"
	"glintfed.org/ent/user"
	"glintfed.org/internal/data"
	"glintfed.org/internal/service/internal"
)

type Service interface {
	SharedInbox(w http.ResponseWriter, r *http.Request)
	UserInbox(w http.ResponseWriter, r *http.Request)
	Webfinger(w http.ResponseWriter, r *http.Request)
	NodeinfoWellKnown(w http.ResponseWriter, r *http.Request)
	HostMeta(w http.ResponseWriter, r *http.Request)
	Nodeinfo(w http.ResponseWriter, r *http.Request)
}

func New(cfg data.Config, client *data.Client) Service {
	return &svc{
		cfg:    cfg,
		client: client,
	}
}

type svc struct {
	cfg    data.Config
	client *data.Client
}

func (s *svc) NodeinfoWellKnown(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.NodeinfoWellKnown")
	defer span.End()

	if !s.cfg.App.Federation.NodeInfo.Enabled {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resp := map[string]any{
		"links": []map[string]string{
			{
				"href": s.cfg.App.URL + "/api/nodeinfo/2.0.json",
				"rel":  "http://nodeinfo.diaspora.software/ns/schema/2.0",
			},
		},
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *svc) Nodeinfo(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.Nodeinfo")
	defer span.End()

	if !s.cfg.App.Federation.NodeInfo.Enabled {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	totalUsers, _ := s.client.Ent.User.Query().Count(r.Context())
	activeMonth, _ := s.client.Ent.User.Query().
		Where(
			user.Or(
				user.UpdatedAtGT(time.Now().Add(-5 * 7 * 24 * time.Hour)),
				user.LastActiveAtGT(time.Now().Add(-5 * 7 * 24 * time.Hour)),
			),
		).Count(r.Context())
	activeHalfyear, _ := s.client.Ent.User.Query().
		Where(
			user.Or(
				user.LastActiveAtGT(time.Now().AddDate(0, -6, 0)),
				user.UpdatedAtGT(time.Now().AddDate(0, -6, 0)),
			),
		).Count(r.Context())

	localPosts, _ := s.client.Ent.Status.Query().
		Where(
			status.LocalEQ(true),
		).Count(r.Context())

	resp := map[string]any{
		"metadata": map[string]any{
			"nodeName": s.cfg.App.Name,
			"software": map[string]string{
				"homepage": "https://glintfed.org",
				"repo":     "https://github.com/glintfed/glintfed",
			},
			"config": map[string]any{
				"features": map[string]any{},
			},
		},
		"protocols": []string{"activitypub"},
		"services": map[string]any{
			"inbound":  []string{},
			"outbound": []string{},
		},
		"software": map[string]string{
			"name":    "glintfed",
			"version": s.cfg.App.Version,
		},
		"usage": map[string]any{
			"localPosts":    localPosts,
			"localComments": 0,
			"users": map[string]int{
				"total":          totalUsers,
				"activeHalfyear": activeHalfyear,
				"activeMonth":    activeMonth,
			},
		},
		"version":           "2.0",
		"openRegistrations": s.cfg.App.Auth.EnableRegistration,
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *svc) SharedInbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.SharedInbox")
	defer span.End()
	// TODO: Implement
}

func (s *svc) UserInbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.UserInbox")
	defer span.End()
	// TODO: Implement
}

func (s *svc) Webfinger(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.Webfinger")
	defer span.End()
	// TODO: Implement
}

func (s *svc) HostMeta(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.HostMeta")
	defer span.End()
	// TODO: Implement
}
