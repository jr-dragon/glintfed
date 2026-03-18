package federation

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/synctest"

	"github.com/go-chi/chi/v5"
	"glintfed.org/internal/data"
	"glintfed.org/internal/usecase/worker"
)

func TestSvc_SharedInbox(t *testing.T) {
	t.Run("Disabled", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.Activitypub.Enabled = false
		s := New(cfg, nil, nil, nil, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/federation/sharedinbox", nil)
		w := httptest.NewRecorder()

		s.SharedInbox(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.Activitypub.Enabled = true
		cfg.App.Federation.Activitypub.SharedInbox = true
		s := New(cfg, nil, nil, nil, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/federation/sharedinbox", bytes.NewBufferString("invalid"))
		w := httptest.NewRecorder()

		s.SharedInbox(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})

	t.Run("BlockedDomain", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.Activitypub.Enabled = true
		cfg.App.Federation.Activitypub.SharedInbox = true

		mockIUC := &InstanceUsecaseMock{
			GetBlockedDomainsFunc: func(ctx context.Context) (map[string]struct{}, error) {
				return map[string]struct{}{"blocked.com": {}}, nil
			},
		}
		s := New(cfg, mockIUC, nil, nil, nil)

		payload := worker.InboxPayload{
			ID: "https://blocked.com/activity/1",
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/api/federation/sharedinbox", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		s.SharedInbox(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})

	t.Run("DeletePerson_Success", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			cfg := data.Config{}
			cfg.App.Federation.Activitypub.Enabled = true
			cfg.App.Federation.Activitypub.SharedInbox = true

			mockIUC := &InstanceUsecaseMock{
				GetBlockedDomainsFunc: func(ctx context.Context) (map[string]struct{}, error) {
					return map[string]struct{}{}, nil
				},
			}
			mockPUC := &ProfileUsecaseMock{
				RemoteUrlExistsFunc: func(ctx context.Context, url string) (bool, error) {
					return true, nil
				},
			}
			mockWUC := &WorkerUsecaseMock{
				DeleteFunc: func(ctx context.Context, header http.Header, payload worker.InboxPayload) error { return nil },
			}
			s := New(cfg, mockIUC, mockPUC, nil, mockWUC)

			typ := "Delete"
			payload := worker.InboxPayload{
				ID:   "https://example.com/activity/1",
				Type: &typ,
				Object: &worker.InboxPayloadObject{
					ID:   "https://example.com/users/alice",
					Type: "Person",
				},
			}
			body, _ := json.Marshal(payload)
			req := httptest.NewRequest(http.MethodPost, "/api/federation/sharedinbox", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			s.SharedInbox(w, req)
			synctest.Wait()

			if w.Code != http.StatusNoContent {
				t.Errorf("expected status 204, got %d", w.Code)
			}
			if len(mockWUC.DeleteCalls()) != 1 {
				t.Error("expected Delete to be called once")
			}
		})

	})

	t.Run("Follow_Success", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			cfg := data.Config{}
			cfg.App.Federation.Activitypub.Enabled = true
			cfg.App.Federation.Activitypub.SharedInbox = true

			mockIUC := &InstanceUsecaseMock{
				GetBlockedDomainsFunc: func(ctx context.Context) (map[string]struct{}, error) {
					return map[string]struct{}{}, nil
				},
			}
			mockWUC := &WorkerUsecaseMock{
				InboxFunc: func(ctx context.Context, header http.Header, payload worker.InboxPayload) error { return nil },
			}
			s := New(cfg, mockIUC, nil, nil, mockWUC)

			typ := "Follow"
			payload := worker.InboxPayload{
				ID:   "https://example.com/activity/1",
				Type: &typ,
			}
			body, _ := json.Marshal(payload)
			req := httptest.NewRequest(http.MethodPost, "/api/federation/sharedinbox", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			s.SharedInbox(w, req)
			synctest.Wait()

			if w.Code != http.StatusNoContent {
				t.Errorf("expected status 204, got %d", w.Code)
			}
			if len(mockWUC.InboxCalls()) != 1 {
				t.Error("expected Inbox to be called once")
			}
		})
	})
}

func TestSvc_UserInbox(t *testing.T) {
	t.Run("Disabled", func(t *testing.T) {
		cfg := data.Config{}
		cfg.App.Federation.Activitypub.Enabled = false
		s := New(cfg, nil, nil, nil, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/federation/users/alice/inbox", nil)
		w := httptest.NewRecorder()

		s.UserInbox(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("DeleteTombstone_Success", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			cfg := data.Config{}
			cfg.App.Federation.Activitypub.Enabled = true
			cfg.App.Federation.Activitypub.Inbox = true

			mockIUC := &InstanceUsecaseMock{
				GetBlockedDomainsFunc: func(ctx context.Context) (map[string]struct{}, error) {
					return map[string]struct{}{}, nil
				},
			}
			mockSUC := &StatusUsecaseMock{
				ObjectUrlExistsFunc: func(ctx context.Context, url string) (bool, error) {
					return true, nil
				},
			}
			mockWUC := &WorkerUsecaseMock{
				DeleteFunc: func(ctx context.Context, header http.Header, payload worker.InboxPayload) error { return nil },
			}
			s := New(cfg, mockIUC, nil, mockSUC, mockWUC)

			typ := "Delete"
			payload := worker.InboxPayload{
				ID:   "https://example.com/activity/1",
				Type: &typ,
				Object: &worker.InboxPayloadObject{
					ID:   "https://example.com/statuses/123",
					Type: "Tombstone",
				},
			}
			body, _ := json.Marshal(payload)
			req := httptest.NewRequest(http.MethodPost, "/api/federation/users/alice/inbox", bytes.NewBuffer(body))

			// Setup chi context for URL param
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("username", "alice")
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()

			s.UserInbox(w, req)
			synctest.Wait()

			if w.Code != http.StatusNoContent {
				t.Errorf("expected status 204, got %d", w.Code)
			}
			if len(mockWUC.DeleteCalls()) != 1 {
				t.Error("expected Delete to be called once")
			}
		})
	})

	t.Run("Follow_Success", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			cfg := data.Config{}
			cfg.App.Federation.Activitypub.Enabled = true
			cfg.App.Federation.Activitypub.Inbox = true

			mockIUC := &InstanceUsecaseMock{
				GetBlockedDomainsFunc: func(ctx context.Context) (map[string]struct{}, error) {
					return map[string]struct{}{}, nil
				},
			}
			mockWUC := &WorkerUsecaseMock{
				ValidateFunc: func(ctx context.Context, username string, header http.Header, payload worker.InboxPayload) error {
					return nil
				},
			}
			s := New(cfg, mockIUC, nil, nil, mockWUC)

			typ := "Follow"
			payload := worker.InboxPayload{
				ID:   "https://example.com/activity/1",
				Type: &typ,
			}
			body, _ := json.Marshal(payload)
			req := httptest.NewRequest(http.MethodPost, "/api/federation/users/alice/inbox", bytes.NewBuffer(body))

			// Setup chi context for URL param
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("username", "alice")
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()

			s.UserInbox(w, req)
			synctest.Wait()

			if w.Code != http.StatusNoContent {
				t.Errorf("expected status 204, got %d", w.Code)
			}
			if len(mockWUC.ValidateCalls()) != 1 {
				t.Error("expected Validate to be called once")
			}
			if mockWUC.ValidateCalls()[0].Username != "alice" {
				t.Errorf("expected username alice, got %s", mockWUC.ValidateCalls()[0].Username)
			}
		})
	})
}
