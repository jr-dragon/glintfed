package instanceactor

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/lib/urls"
	"glintfed.org/internal/service/internal"
)

type Service interface {
	Profile(w http.ResponseWriter, r *http.Request)
	Inbox(w http.ResponseWriter, r *http.Request)
	Outbox(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_instance_actor_getter.go . InstanceActorGetter
type InstanceActorGetter interface {
	Get(ctx context.Context) (*ent.InstanceActor, error)
}

func New(cfg *data.Config, iag InstanceActorGetter) Service {
	return &svc{
		cfg: cfg,

		iag: iag,
	}
}

type svc struct {
	cfg *data.Config

	iag InstanceActorGetter
}

func (s *svc) Profile(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Profile")
	defer span.End()

	parsed, err := url.Parse(s.cfg.App.Url)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to parse app url", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	ia, err := s.iag.Get(r.Context())
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to get instance actor", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(ia.GetActor(parsed.String(), parsed.Host)); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode instance actor", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
}

func (s *svc) Inbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Inbox")
	defer span.End()

	w.WriteHeader(http.StatusNoContent)
}

func (s *svc) Outbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Outbox")
	defer span.End()

	res := map[string]any{
		"@context": []any{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
			map[string]any{
				"manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
				"toot":                      "http://joinmastodon.org/ns#",
				"featured": map[string]string{
					"@id":   "toot:featured",
					"@type": "@id",
				},
				"featuredTags": map[string]string{
					"@id":   "toot:featuredTags",
					"@type": "@id",
				},
				"alsoKnownAs": map[string]string{
					"@id":   "as:alsoKnownAs",
					"@type": "@id",
				},
				"movedTo": map[string]string{
					"@id":   "as:movedTo",
					"@type": "@id",
				},
				"schema":           "http://schema.org#",
				"PropertyValue":    "schema:PropertyValue",
				"value":            "schema:value",
				"discoverable":     "toot:discoverable",
				"Device":           "toot:Device",
				"Ed25519Signature": "toot:Ed25519Signature",
				"Ed25519Key":       "toot:Ed25519Key",
				"Curve25519Key":    "toot:Curve25519Key",
				"EncryptedMessage": "toot:EncryptedMessage",
				"publicKeyBase64":  "toot:publicKeyBase64",
				"deviceId":         "toot:deviceId",
				"claim": map[string]string{
					"@type": "@id",
					"@id":   "toot:claim",
				},
				"fingerprintKey": map[string]string{
					"@type": "@id",
					"@id":   "toot:fingerprintKey",
				},
				"identityKey": map[string]string{
					"@type": "@id",
					"@id":   "toot:identityKey",
				},
				"devices": map[string]string{
					"@type": "@id",
					"@id":   "toot:devices",
				},
				"messageFranking": "toot:messageFranking",
				"messageType":     "toot:messageType",
				"cipherText":      "toot:cipherText",
				"suspended":       "toot:suspended",
			},
		},
		"id":         urls.MustJoinPath(s.cfg.App.Url, "/i/actor/outbot"),
		"type":       "OrderedCollection",
		"totalItems": 0,
		"first":      urls.MustJoinPath(s.cfg.App.Url, "/i/actor/outbox?page=true"),
		"last":       urls.MustJoinPath(s.cfg.App.Url, "/i/actor/outbox?min_id=0&page=true"),
	}

	w.Header().Set("Content-Type", "application/activity+json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
