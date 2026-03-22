package ent

import (
	"log/slog"
	"net/url"
	"strings"
)

const ProfileBase = "/i/actor"

func (ia *InstanceActor) Permalink(baseUrl string, suffixes ...string) string {
	res, err := url.JoinPath(baseUrl, ProfileBase, strings.Join(suffixes, ""))
	if err != nil {
		slog.Error("failed to join path",
			slog.String("baseUrl", baseUrl),
			slog.String("profile_base", ProfileBase),
			slog.Any("suffixes", suffixes),
		)
	}

	return res
}

func (ia *InstanceActor) GetActor(url, domain string) map[string]any {
	return map[string]any{
		"@context": []any{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
			map[string]any{
				"manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
				"toot":                      "http://joinmastodon.org/ns#",
				"featured": map[string]any{
					"@id":   "toot:featured",
					"@type": "@id",
				},
				"featuredTags": map[string]any{
					"@id":   "toot:featuredTags",
					"@type": "@id",
				},
				"alsoKnownAs": map[string]any{
					"@id":   "as:alsoKnownAs",
					"@type": "@id",
				},
				"movedTo": map[string]any{
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
				"claim": map[string]any{
					"@type": "@id",
					"@id":   "toot:claim",
				},
				"fingerprintKey": map[string]any{
					"@type": "@id",
					"@id":   "toot:fingerprintKey",
				},
				"identityKey": map[string]any{
					"@type": "@id",
					"@id":   "toot:identityKey",
				},
				"devices": map[string]any{
					"@type": "@id",
					"@id":   "toot:devices",
				},
				"messageFranking": "toot:messageFranking",
				"messageType":     "toot:messageType",
				"cipherText":      "toot:cipherText",
				"suspended":       "toot:suspended",
			},
		},
		"id":                ia.Permalink(url),
		"type":              "Application",
		"inbox":             ia.Permalink(url, "/inbox"),
		"outbox":            ia.Permalink(url, "/outbox"),
		"preferredUsername": domain,
		"publicKey": map[string]any{
			"id":           ia.Permalink(url, "#main-key"),
			"owner":        ia.Permalink(url),
			"publicKeyPem": ia.PublicKey,
		},
		"manuallyApprovesFollowers": true,
		"url":                       url + "/site/kb/instance-actor",
	}
}
