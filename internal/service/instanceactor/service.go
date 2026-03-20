package instanceactor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"glintfed.org/internal/data"
	"glintfed.org/internal/service/internal"
)

type Service interface {
	Profile(w http.ResponseWriter, r *http.Request)
	Inbox(w http.ResponseWriter, r *http.Request)
	Outbox(w http.ResponseWriter, r *http.Request)
}

func New(cfg *data.Config) Service {
	return &svc{
		cfg: cfg,
	}
}

type svc struct {
	cfg *data.Config
}

func (s *svc) Profile(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Profile")
	defer span.End()
	// TODO: Implement
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
				"schema":        "http://schema.org#",
				"PropertyValue": "schema:PropertyValue",
				"value":         "schema:value",
				"discoverable":  "toot:discoverable",
				"Device":        "toot:Device",
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
		"id":         fmt.Sprintf("%s/i/actor/outbox", s.cfg.App.Url),
		"type":       "OrderedCollection",
		"totalItems": 0,
		"first":      fmt.Sprintf("%s/i/actor/outbox?page=true", s.cfg.App.Url),
		"last":       fmt.Sprintf("%s/i/actor/outbox?min_id=0&page=true", s.cfg.App.Url),
	}

	w.Header().Set("Content-Type", "application/activity+json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
