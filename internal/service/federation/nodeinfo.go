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
	Features Features `json:"features"`
}

type Features struct {
	Version             string             `json:"version"`
	EnableRegistration  bool               `json:"open_registration"`
	ShowLegalNoticeLink bool               `json:"show_legal_notice_link"`
	Uploader            UploaderFeature      `json:"uploader"`
	Activitypub         ActivitypubFeature `json:"activitypub"`
	AB                  map[string]bool    `json:"ab"`
	Site                SiteFeature        `json:"site"`
	Account             AccountFeature     `json:"account"`
	Username            UsernameFeature    `json:"username"`
	Features            FeaturesFeature    `json:"features"`
}

type UploaderFeature struct {
	MaxPhotoSize        int      `json:"max_photo_size"`
	MaxCaptionLength    int      `json:"max_caption_length"`
	MaxAltextLength     int      `json:"max_altext_length"`
	AlbumLimit          int      `json:"album_limit"`
	ImageQuality        int      `json:"image_quality"`
	MaxCollectionLength int      `json:"max_collection_length"`
	OptimizeImage       bool     `json:"optimize_image"`
	OptimizeVideo       bool     `json:"optimize_video"`
	MediaTypes          string   `json:"media_types"`
	MimeTypes           []string `json:"mime_types"`
	EnforceAcountLimit  bool     `json:"enforce_account_limit"`
}

type ActivitypubFeature struct {
	Enabled      bool `json:"enabled"`
	RemoteFollow bool `json:"remote_follow"`
}

type SiteFeature struct {
	Name        string `json:"name"`
	Domain      string `json:"domain"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type AccountFeature struct {
	MaxAvatarSize     int `json:"max_avatar_size"`
	MaxBioLength      int `json:"max_bio_length"`
	MaxNameLength     int `json:"max_name_legth"`
	MinPasswordLength int `json:"min_password_length"`
	MaxAccountSize    int `json:"max_account_size"`
}

type UsernameFeature struct {
	Remote RemoteUsernameFeature `json:"remote"`
}

type FeaturesFeature struct {
	Timelines          TimelineFeature  `json:"timelines"`
	MobileAPIs         bool             `json:"mobile_apis"`
	MobileRegistration bool             `json:"mobile_registration"`
	Stories            bool             `json:"stories"`
	Video              bool             `json:"video"`
	Import             ImportFeature    `json:"import"`
	Label              map[string]Label `json:"label"`
	HLS                HLSFeature       `json:"hls"`
	Groups             bool             `json:"groups"`
}

type TimelineFeature struct {
	Local   bool `json:"local"`
	Network bool `json:"network"`
}

type ImportFeature struct {
	Instagram bool `json:"instagram"`
	Mastodon  bool `json:"mastodon"`
	Pixelfed  bool `json:"pixelfed"`
}

type Label struct {
	Enabled bool   `json:"enabled"`
	Org     string `json:"org"`
	Url     string `json:"url"`
}

type HLSFeature struct {
	Enabled  bool   `json:"enabled"`
	Debug    bool   `json:"debug"`
	P2P      bool   `json:"p2p"`
	P2PDebug bool   `json:"p2p_debug"`
	Tracker  string `json:"tracker"`
	Ice      string `json:"ice"`
}

type RemoteUsernameFeature struct {
	Formats []string `json:"formats"`
	Format  bool     `json:"format"`
	Custom  string   `json:"custom"`
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
				Features: s.getFeatures(),
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

func (s *svc) getFeatures() Features {
	return Features{
		Version:            s.cfg.App.Version,
		EnableRegistration: s.cfg.App.Auth.EnableRegistration,
	}
}
