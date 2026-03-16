package federation

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

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

	if err := json.NewDecoder(bytes.NewBuffer(raw)).Decode(&payload); err != nil {
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
		case "Follow", "Accept":
			s.wuc.Validate(r.Context(), r.URL.Query().Get("username"), r.Header, payload)
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
			fallthrough
		default:
			slog.ErrorContext(r.Context(), "invalid payload", slog.Any("type", payload))
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	} else {
		s.wuc.Inbox(r.Context(), r.Header, payload)
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
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
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
		case "Follow", "Accept":
			s.wuc.Validate(r.Context(), r.URL.Query().Get("username"), r.Header, payload)
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
			fallthrough
		default:
			slog.ErrorContext(r.Context(), "invalid payload", slog.Any("type", payload))
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	} else {
		s.wuc.Validate(r.Context(), r.URL.Query().Get("username"), r.Header, payload)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *svc) validInboxDomain(ctx context.Context, domain string) error {
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

func (s *svc) handleDelete(ctx context.Context, header http.Header, payload worker.InboxPayload) error {
	switch payload.Object.Type {
	case "Person":
		if exists, err := s.puc.RemoteUrlExists(ctx, payload.Object.ID); err != nil {
			return err
		} else if !exists {
			return ErrMissingUrl
		}
	case "Tombstone":
		if exists, err := s.suc.ObjectUrlExists(ctx, payload.Object.ID); err != nil {
			return err
		} else if !exists {
			return ErrMissingUrl
		}
	case "Story":
		break
	default:
		return ErrInvalidType
	}

	s.wuc.Delete(ctx, header, payload)
	return nil
}
