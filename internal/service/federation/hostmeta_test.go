package federation

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"

	"glintfed.org/internal/data"
)

func TestSvc_HostMeta(t *testing.T) {
	t.Run("Disabled", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.Webfinger.Enabled = false
		s := New(cfg, nil, nil, nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/.well-known/host-meta", nil)
		w := httptest.NewRecorder()

		s.HostMeta(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("Enabled", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.Webfinger.Enabled = true
		cfg.App.Url = "https://example.com"
		s := New(cfg, nil, nil, nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/.well-known/host-meta", nil)
		w := httptest.NewRecorder()

		s.HostMeta(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if contentType := w.Header().Get("Content-Type"); contentType != "application/xrd+xml" {
			t.Errorf("expected content-type application/xrd+xml, got %s", contentType)
		}

		var resp HostMetaXRD
		if err := xml.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatalf("failed to decode xml: %v", err)
		}

		if len(resp.Links) != 1 {
			t.Fatalf("expected 1 link, got %d", len(resp.Links))
		}

		expectedTemplate := "https://example.com/.well-known/webfinger?resource={uri}"
		if resp.Links[0].Template != expectedTemplate {
			t.Errorf("expected template %s, got %s", expectedTemplate, resp.Links[0].Template)
		}
	})
}
