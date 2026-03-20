package federation

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

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
				Href: s.cfg.App.Url + "/api/nodeinfo/2.0.json",
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
	Uploader            UploaderFeature    `json:"uploader"`
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
	Format  string   `json:"format"`
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

	totalUsers, err := s.um.GetTotalUsers(r.Context())
	if err != nil {
		const msg = "failed to get total users"
		slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
		http.Error(w, msg, http.StatusInternalServerError)
	}

	activeMonth, err := s.um.GetMonthActiveUsers(r.Context())
	if err != nil {
		const msg = "failed to get month active users"
		slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
		http.Error(w, msg, http.StatusInternalServerError)
	}

	activeHalfyear, err := s.um.GetHalfYearActiveUsers(r.Context())
	if err != nil {
		const msg = "failed to get half year active users"
		slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	localPosts, err := s.sm.GetLocalPostsCount(r.Context())
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
	u, err := url.Parse(s.cfg.App.Url)
	if err != nil {
		slog.Warn("failed to parse app url", logs.ErrAttr(err))
	}

	var domain string
	if u != nil {
		domain = u.Host
	}

	return Features{
		Version:             s.cfg.App.Version,
		EnableRegistration:  s.cfg.App.Auth.EnableRegistration,
		ShowLegalNoticeLink: s.cfg.App.Instance.HasLegalNotice,
		Uploader: UploaderFeature{
			MaxPhotoSize:        s.cfg.App.MaxPhotoSize,
			MaxCaptionLength:    s.cfg.App.MaxCaptionLength,
			MaxAltextLength:     s.cfg.App.MaxAltextLength,
			AlbumLimit:          s.cfg.App.MaxAlbumLength,
			ImageQuality:        s.cfg.App.ImageQuality,
			MaxCollectionLength: s.cfg.App.MaxCollectionLength,
			OptimizeImage:       s.cfg.App.OptimizeImage,
			OptimizeVideo:       s.cfg.App.OptimizeVideo,
			MediaTypes:          s.cfg.App.MediaTypes,
			MimeTypes:           strings.Split(s.cfg.App.MediaTypes, ","),
			EnforceAcountLimit:  s.cfg.App.EnforceAcountLimit,
		},
		Activitypub: ActivitypubFeature{
			Enabled:      s.cfg.App.Federation.Activitypub.Enabled,
			RemoteFollow: s.cfg.App.Federation.Activitypub.RemoteFollow,
		},
		Site: SiteFeature{
			Name:        s.cfg.App.Name,
			Domain:      domain,
			Url:         s.cfg.App.Url,
			Description: s.cfg.App.Description,
		},
		Account: AccountFeature{
			MaxAvatarSize:     s.cfg.App.MaxAvatarSize,
			MaxBioLength:      s.cfg.App.MaxBioLength,
			MaxNameLength:     s.cfg.App.MaxNameLength,
			MinPasswordLength: s.cfg.App.MinPasswordLength,
			MaxAccountSize:    s.cfg.App.MaxAccountSize,
		},
		Username: UsernameFeature{
			Remote: RemoteUsernameFeature{
				Formats: s.cfg.App.Instance.Username.Remote.Formats,
				Format:  s.cfg.App.Instance.Username.Remote.Format,
				Custom:  s.cfg.App.Instance.Username.Remote.Custom,
			},
		},
		Features: FeaturesFeature{
			Timelines: TimelineFeature{
				Local:   true,
				Network: s.cfg.App.Federation.NetworkTimeline,
			},
			MobileAPIs:         s.cfg.App.Auth.EnableOAuth,
			MobileRegistration: s.cfg.App.Auth.InAppRegistration,
			Stories:            s.cfg.App.Instance.Stories.Enabled,
			Video:              strings.Contains(s.cfg.App.MediaTypes, "video/mp4"),
			Import: ImportFeature{
				Instagram: s.cfg.App.Import.Instagram.Enabled,
				Mastodon:  false,
				Pixelfed:  false,
			},
			Label: map[string]Label{
				"covid": Label{
					Enabled: s.cfg.App.Instance.Label.Covid.Enabled,
					Org:     s.cfg.App.Instance.Label.Covid.Org,
					Url:     s.cfg.App.Instance.Label.Covid.Url,
				},
			},
			HLS: HLSFeature{
				Enabled:  s.cfg.App.Media.HLS.Enabled,
				Debug:    s.cfg.App.Media.HLS.Debug,
				P2P:      s.cfg.App.Media.HLS.P2P,
				P2PDebug: s.cfg.App.Media.HLS.P2PDebug,
				Tracker:  s.cfg.App.Media.HLS.Tracker,
				Ice:      s.cfg.App.Media.HLS.Ice,
			},
			Groups: s.cfg.App.Groups.Enabled,
		},
	}
}
