package federation

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

type sharedInboxPayload struct {
	ID     string                    `json:"id"`
	Type   *string                   `json:"type,omitempty"`
	Object *sharedInboxPayloadObject `json:"object,omitempty"`
}

type sharedInboxPayloadObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (s *svc) SharedInbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.SharedInbox")
	defer span.End()

	if !s.cfg.App.Federation.Activitypub.Enabled ||
		!s.cfg.App.Federation.Activitypub.SharedInbox {
		http.NotFound(w, r)
		return
	}

	var payload sharedInboxPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		slog.ErrorContext(r.Context(), "failed to decode json payload", logs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := s.validSharedInboxDomain(r.Context(), payload.ID); err != nil {
		slog.ErrorContext(r.Context(), "invalid domain", logs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if payload.Type != nil {
		if *payload.Type == "Delete" && payload.Object != nil {
			switch payload.Object.Type {
			case "Person":
				if exists, err := s.puc.RemoteUrlExists(r.Context(), payload.Object.ID); err == nil && exists {
					// TODO: go worker.DeleteInbox(r.Context(), r.Header, payload)
				} else if err != nil {
					const msg = "failed to check remote url exists"
					slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
					http.Error(w, msg, http.StatusInternalServerError)
					return
				}
			case "Tombstone":
				if exists, err := s.suc.ObjectUrlExists(r.Context(), payload.Object.ID); err == nil && exists {
					// todo go worker.Delete(r.Context(), r.Header, payload)
				} else if err != nil {
					const msg = "failed to check object url exists"
					slog.ErrorContext(r.Context(), msg, logs.ErrAttr(err))
					http.Error(w, msg, http.StatusInternalServerError)
				}
			case "Story":
				// TODO: go worker.DeleteStory(r.Context(), r.Header, payload)
			}
		} else if *payload.Type == "Follow" || *payload.Type == "Accept" {
			// TODO: go worker.InboxFollow(r.Context(), r.Header, payload)
		} else {
			const msg = "invalid payload"
			slog.ErrorContext(r.Context(), msg, slog.Any("type", payload))
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	} else {
		// TODO: go inboxworker.Sahred(r.Context(), r.Header, payload)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *svc) validSharedInboxDomain(ctx context.Context, domain string) error {
	parsed, err := url.Parse(domain)
	if err != nil {
		return err
	}

	blocked, err := s.iuc.GetBlockedDomains(ctx)
	if err != nil {
		return err
	}

	if _, ok := blocked[parsed.Host]; !ok {
		return fmt.Errorf("host %s has been blocked", parsed.Host)
	}

	return nil
}
