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
		cfg := &data.Config{}
		cfg.App.Federation.NodeInfo.Enabled = false
		s := New(cfg, &InstanceUsecaseMock{}, nil, nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/.well-known/nodeinfo", nil)
		w := httptest.NewRecorder()

		s.NodeinfoWellKnown(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("Enabled", func(t *testing.T) {
		cfg := &data.Config{}
		cfg.App.Federation.NodeInfo.Enabled = true
		cfg.App.Url = "https://example.com"
		s := New(cfg, &InstanceUsecaseMock{}, nil, nil, nil)

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
		cfg := &data.Config{}
		cfg.App.Federation.NodeInfo.Enabled = false
		s := New(cfg, &InstanceUsecaseMock{}, nil, nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/nodeinfo/2.0.json", nil)
		w := httptest.NewRecorder()

		s.Nodeinfo(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("Enabled", func(t *testing.T) {
		cfg := &data.Config{}
		cfg.App.Federation.NodeInfo.Enabled = true
		cfg.App.Name = "TestNode"
		cfg.App.Version = "1.0.0"
		cfg.App.Url = "https://example.com"
		cfg.App.Description = "A test node"
		cfg.App.Auth.EnableRegistration = true
		cfg.App.Instance.HasLegalNotice = true
		cfg.App.MaxPhotoSize = 1024
		cfg.App.MaxCaptionLength = 500
		cfg.App.Federation.Activitypub.Enabled = true
		cfg.App.MaxAvatarSize = 2048
		cfg.App.MaxBioLength = 150
		cfg.App.Groups.Enabled = true

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
		s := New(cfg, mockIUC, nil, nil, nil)

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

		// Test Features
		features := resp.Metadata.Config.Features
		if features.Version != "1.0.0" {
			t.Errorf("expected features version 1.0.0, got %s", features.Version)
		}
		if features.EnableRegistration != true {
			t.Errorf("expected features open_registration true, got %v", features.EnableRegistration)
		}
		if features.ShowLegalNoticeLink != true {
			t.Errorf("expected show_legal_notice_link true, got %v", features.ShowLegalNoticeLink)
		}
		if features.Uploader.MaxPhotoSize != 1024 {
			t.Errorf("expected max_photo_size 1024, got %d", features.Uploader.MaxPhotoSize)
		}
		if features.Uploader.MaxCaptionLength != 500 {
			t.Errorf("expected max_caption_length 500, got %d", features.Uploader.MaxCaptionLength)
		}
		if features.Activitypub.Enabled != true {
			t.Errorf("expected activitypub enabled true, got %v", features.Activitypub.Enabled)
		}
		if features.Site.Name != "TestNode" {
			t.Errorf("expected site name TestNode, got %s", features.Site.Name)
		}
		if features.Site.Url != "https://example.com" {
			t.Errorf("expected site url https://example.com, got %s", features.Site.Url)
		}
		if features.Site.Description != "A test node" {
			t.Errorf("expected site description A test node, got %s", features.Site.Description)
		}
		if features.Account.MaxAvatarSize != 2048 {
			t.Errorf("expected max_avatar_size 2048, got %d", features.Account.MaxAvatarSize)
		}
		if features.Account.MaxBioLength != 150 {
			t.Errorf("expected max_bio_length 150, got %d", features.Account.MaxBioLength)
		}
		if features.Features.Groups != true {
			t.Errorf("expected groups enabled true, got %v", features.Features.Groups)
		}
	})
}
