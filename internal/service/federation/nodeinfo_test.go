package federation

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"glintfed.org/internal/data"
)

func TestSvc_NodeinfoWellKnown(t *testing.T) {
	t.Run("Disabled", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.NodeInfo.Enabled = false
		s := New(cfg, &InstanceUsecaseMock{})

		req := httptest.NewRequest(http.MethodGet, "/.well-known/nodeinfo", nil)
		w := httptest.NewRecorder()

		s.NodeinfoWellKnown(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("Enabled", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.NodeInfo.Enabled = true
		cfg.App.URL = "https://example.com"
		s := New(cfg, &InstanceUsecaseMock{})

		req := httptest.NewRequest(http.MethodGet, "/.well-known/nodeinfo", nil)
		w := httptest.NewRecorder()

		s.NodeinfoWellKnown(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		var resp NodeInfoWellKnownResponse
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatal(err)
		}

		if len(resp.Links) != 1 {
			t.Fatalf("expected 1 link, got %d", len(resp.Links))
		}

		expectedHref := "https://example.com/api/nodeinfo/2.0.json"
		if resp.Links[0].Href != expectedHref {
			t.Errorf("expected href %s, got %s", expectedHref, resp.Links[0].Href)
		}
	})
}

func TestSvc_Nodeinfo(t *testing.T) {
	t.Run("Disabled", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.NodeInfo.Enabled = false
		s := New(cfg, &InstanceUsecaseMock{})

		req := httptest.NewRequest(http.MethodGet, "/api/nodeinfo/2.0.json", nil)
		w := httptest.NewRecorder()

		s.Nodeinfo(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("Enabled", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.NodeInfo.Enabled = true
		cfg.App.Name = "TestNode"
		cfg.App.Version = "1.0.0"
		cfg.App.Auth.EnableRegistration = true

		mockIUC := &InstanceUsecaseMock{
			GetTotalUsersFunc: func(ctx context.Context) (int, error) {
				return 1, nil
			},
			GetMonthActiveUsersFunc: func(ctx context.Context) (int, error) {
				return 1, nil
			},
			GetHalfYearActiveUsersFunc: func(ctx context.Context) (int, error) {
				return 1, nil
			},
			GetLocalPostsCountFunc: func(ctx context.Context) (int, error) {
				return 1, nil
			},
		}
		s := New(cfg, mockIUC)

		req := httptest.NewRequest(http.MethodGet, "/api/nodeinfo/2.0.json", nil)
		w := httptest.NewRecorder()

		s.Nodeinfo(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		var resp NodeInfoResponse
		if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
			t.Fatal(err)
		}

		if resp.Metadata.NodeName != "TestNode" {
			t.Errorf("expected nodeName TestNode, got %s", resp.Metadata.NodeName)
		}

		if resp.Software.Version != "1.0.0" {
			t.Errorf("expected version 1.0.0, got %s", resp.Software.Version)
		}

		if resp.Usage.Users.Total != 1 {
			t.Errorf("expected total users 1, got %d", resp.Usage.Users.Total)
		}

		if resp.Usage.LocalPosts != 1 {
			t.Errorf("expected local posts 1, got %d", resp.Usage.LocalPosts)
		}

		if resp.OpenRegistrations != true {
			t.Errorf("expected openRegistrations true, got %v", resp.OpenRegistrations)
		}
	})
}
