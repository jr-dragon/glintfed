package federation

import (
	"encoding/xml"
	"log/slog"
	"net/http"
	"net/url"

	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

type HostMetaXRD struct {
	XMLName xml.Name       `xml:"http://docs.oasis-open.org/ns/xri/xrd-1.0 XRD"`
	Links   []HostMetaLink `xml:"Link"`
}

type HostMetaLink struct {
	Rel      string `xml:"rel,attr"`
	Type     string `xml:"type,attr"`
	Template string `xml:"template,attr"`
}

func (s *svc) HostMeta(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.HostMeta")
	defer span.End()

	if !s.cfg.App.Federation.Webfinger.Enabled {
		http.NotFound(w, r)
		return
	}

	webfingerPath, err := url.JoinPath(s.cfg.App.Url, "/.well-known/webfinger")
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to join url path", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	xrd := HostMetaXRD{
		Links: []HostMetaLink{
			{
				Rel:      "lrdd",
				Type:     "application/xrd+xml",
				Template: webfingerPath + "?resource={uri}",
			},
		},
	}

	w.Header().Set("Content-Type", "application/xrd+xml")
	w.Write([]byte(xml.Header))
	if err := xml.NewEncoder(w).Encode(xrd); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode host-meta xml", logs.ErrAttr(err))
	}
}
