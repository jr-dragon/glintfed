package federation

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

type NodeInfoWellKnownResponse struct {
	Links []NodeInfoLink `json:"links"`
}

type NodeInfoLink struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
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

	resp := NodeInfoWellKnownResponse{
		Links: []NodeInfoLink{
			{
				Href: s.cfg.App.URL + "/api/nodeinfo/2.0.json",
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
			},
		},
	}

	json.NewEncoder(w).Encode(resp)
}

type NodeInfoResponse struct {
	Metadata          NodeInfoMetadata `json:"metadata"`
	Protocols         []string         `json:"protocols"`
	Services          NodeInfoServices `json:"services"`
	Software          NodeInfoSoftware `json:"software"`
	Usage             NodeInfoUsage    `json:"usage"`
	Version           string           `json:"version"`
	OpenRegistrations bool             `json:"openRegistrations"`
}

type NodeInfoMetadata struct {
	NodeName string           `json:"nodeName"`
	Software MetadataSoftware `json:"software"`
	Config   MetadataConfig   `json:"config"`
}

type MetadataSoftware struct {
	Homepage string `json:"homepage"`
	Repo     string `json:"repo"`
}

type MetadataConfig struct {
	Features map[string]any `json:"features"`
}

type NodeInfoServices struct {
	Inbound  []string `json:"inbound"`
	Outbound []string `json:"outbound"`
}

type NodeInfoSoftware struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type NodeInfoUsage struct {
	LocalPosts    int           `json:"localPosts"`
	LocalComments int           `json:"localComments"`
	Users         NodeInfoUsers `json:"users"`
}

type NodeInfoUsers struct {
	Total          int `json:"total"`
	ActiveHalfyear int `json:"activeHalfyear"`
	ActiveMonth    int `json:"activeMonth"`
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

	totalUsers, err := s.iuc.GetTotalUsers(r.Context())
	if err != nil {
		const msg = "failed to get total users"
		slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
		http.Error(w, msg, http.StatusInternalServerError)
	}

	activeMonth, err := s.iuc.GetMonthActiveUsers(r.Context())
	if err != nil {
		const msg = "failed to get month active users"
		slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
		http.Error(w, msg, http.StatusInternalServerError)
	}

	activeHalfyear, err := s.iuc.GetHalfYearActiveUsers(r.Context())
	if err != nil {
		const msg = "failed to get half year active users"
		slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	localPosts, err := s.iuc.GetLocalPostsCount(r.Context())
	if err != nil {
		const msg = "failed to get local posts count"
		slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	resp := NodeInfoResponse{
		Metadata: NodeInfoMetadata{
			NodeName: s.cfg.App.Name,
			Software: MetadataSoftware{
				Homepage: "https://glintfed.org",
				Repo:     "https://github.com/glintfed/glintfed",
			},
			Config: MetadataConfig{
				Features: map[string]any{},
			},
		},
		Protocols: []string{"activitypub"},
		Services: NodeInfoServices{
			Inbound:  []string{},
			Outbound: []string{},
		},
		Software: NodeInfoSoftware{
			Name:    "glintfed",
			Version: s.cfg.App.Version,
		},
		Usage: NodeInfoUsage{
			LocalPosts:    localPosts,
			LocalComments: 0,
			Users: NodeInfoUsers{
				Total:          totalUsers,
				ActiveHalfyear: activeHalfyear,
				ActiveMonth:    activeMonth,
			},
		},
		Version:           "2.0",
		OpenRegistrations: s.cfg.App.Auth.EnableRegistration,
	}

	json.NewEncoder(w).Encode(resp)
}
