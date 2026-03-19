package federation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
	"glintfed.org/internal/usecase/worker"
)

var (
	ErrMissingUrl  = errors.New("missing url")
	ErrInvalidType = errors.New("invalid type")
)

func (s *svc) SharedInbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.SharedInbox")
	defer span.End()

	if !s.cfg.App.Federation.Activitypub.Enabled ||
		!s.cfg.App.Federation.Activitypub.SharedInbox {
		http.NotFound(w, r)
		return
	}

	var payload worker.InboxPayload
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to read request body", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	payload.Raw = string(raw)

	if err := json.Unmarshal(raw, &payload); err != nil {
		slog.ErrorContext(r.Context(), "failed to decode json payload", logs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := s.validInboxDomain(r.Context(), payload.ID); err != nil {
		slog.ErrorContext(r.Context(), "invalid domain", logs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if payload.Type != nil {
		switch *payload.Type {
		case "Delete":
			if payload.Object != nil {
				if err := s.handleDelete(r.Context(), r.Header, payload); err != nil {
					slog.ErrorContext(r.Context(), "failed to handle delete", logs.ErrAttr(err), slog.Any("payload", payload))
					switch {
					case errors.Is(err, ErrMissingUrl):
						w.WriteHeader(http.StatusNoContent)
					case errors.Is(err, ErrInvalidType):
						w.WriteHeader(http.StatusBadRequest)
					default:
						w.WriteHeader(http.StatusInternalServerError)
					}
				} else {
					w.WriteHeader(http.StatusNoContent)
				}
				return
			}
			slog.ErrorContext(r.Context(), "invalid payload", slog.Any("payload", payload))
			http.Error(w, "", http.StatusBadRequest)
			return
		case "Follow", "Accept":
			h := r.Header.Clone()
			go func() {
				if err := s.wuc.Inbox(r.Context(), h, payload); err != nil {
					slog.ErrorContext(r.Context(), "failed to handle inbox", logs.ErrAttr(err))
				}
			}()
		default:
			h := r.Header.Clone()
			go func() {
				if err := s.wuc.Inbox(r.Context(), h, payload); err != nil {
					slog.ErrorContext(r.Context(), "failed to handle inbox", logs.ErrAttr(err))
				}
			}()
		}
	} else {
		h := r.Header.Clone()
		go func() {
			if err := s.wuc.Inbox(r.Context(), h, payload); err != nil {
				slog.ErrorContext(r.Context(), "failed to handle inbox", logs.ErrAttr(err))
			}
		}()
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *svc) UserInbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.UserInbox")
	defer span.End()

	if !s.cfg.App.Federation.Activitypub.Enabled ||
		!s.cfg.App.Federation.Activitypub.Inbox {
		http.NotFound(w, r)
		return
	}

	var payload worker.InboxPayload
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to read request body", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	payload.Raw = string(raw)

	if err := json.Unmarshal(raw, &payload); err != nil {
		slog.ErrorContext(r.Context(), "failed to decode json payload", logs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := s.validInboxDomain(r.Context(), payload.ID); err != nil {
		slog.ErrorContext(r.Context(), "invalid domain", logs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	username := chi.URLParam(r, "username")

	if payload.Type != nil {
		switch *payload.Type {
		case "Delete":
			if payload.Object != nil {
				if err := s.handleDelete(r.Context(), r.Header, payload); err != nil {
					slog.ErrorContext(r.Context(), "failed to handle delete", logs.ErrAttr(err), slog.Any("payload", payload))
					switch {
					case errors.Is(err, ErrMissingUrl):
						w.WriteHeader(http.StatusNoContent)
					case errors.Is(err, ErrInvalidType):
						w.WriteHeader(http.StatusBadRequest)
					default:
						w.WriteHeader(http.StatusInternalServerError)
					}
				} else {
					w.WriteHeader(http.StatusNoContent)
				}
				return
			}
			slog.ErrorContext(r.Context(), "invalid payload", slog.Any("type", payload))
			http.Error(w, "", http.StatusBadRequest)
			return
		case "Follow", "Accept":
			h := r.Header.Clone()
			go func() {
				if err := s.wuc.Validate(r.Context(), username, h, payload); err != nil {
					slog.ErrorContext(r.Context(), "failed to handle validate inbox", logs.ErrAttr(err))
				}
			}()
		default:
			h := r.Header.Clone()
			go func() {
				if err := s.wuc.Validate(r.Context(), username, h, payload); err != nil {
					slog.ErrorContext(r.Context(), "failed to handle validate inbox", logs.ErrAttr(err))
				}
			}()
		}
	} else {
		h := r.Header.Clone()
		go func() {
			if err := s.wuc.Validate(r.Context(), username, h, payload); err != nil {
				slog.ErrorContext(r.Context(), "failed to handle validate inbox", logs.ErrAttr(err))
			}
		}()
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *svc) validInboxDomain(ctx context.Context, domain string) error {
	parsed, err := url.Parse(domain)
	if err != nil {
		return err
	}

	blocked, err := s.im.GetBlockedDomains(ctx)
	if err != nil {
		return err
	}

	if _, ok := blocked[parsed.Host]; ok {
		return fmt.Errorf("host %s has been blocked", parsed.Host)
	}

	return nil
}

func (s *svc) handleDelete(ctx context.Context, header http.Header, payload worker.InboxPayload) error {
	switch payload.Object.Type {
	case "Person":
		if exists, err := s.pm.RemoteUrlExists(ctx, payload.Object.ID); err != nil {
			return err
		} else if !exists {
			return ErrMissingUrl
		}
	case "Tombstone":
		if exists, err := s.sm.ObjectUrlExists(ctx, payload.Object.ID); err != nil {
			return err
		} else if !exists {
			return ErrMissingUrl
		}
	case "Story":
	default:
		return ErrInvalidType
	}

	h := header.Clone()
	go func() {
		if err := s.wuc.Delete(ctx, h, payload); err != nil {
			slog.ErrorContext(ctx, "failed to handle delete inbox", logs.ErrAttr(err))
		}
	}()

	return nil
}
