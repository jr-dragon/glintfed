package ent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceActor_Permalink(t *testing.T) {
	ia := &InstanceActor{}
	appURL := "https://example.com"

	assert.Equal(t, "https://example.com/i/actor", ia.Permalink(appURL))
	assert.Equal(t, "https://example.com/i/actor/inbox", ia.Permalink(appURL, "/inbox"))
}

func TestInstanceActor_GetActor(t *testing.T) {
	ia := &InstanceActor{
		PublicKey: "test-public-key",
	}
	appURL := "https://example.com"
	appDomain := "example.com"

	actor := ia.GetActor(appURL, appDomain)

	assert.Equal(t, "https://example.com/i/actor", actor["id"])
	assert.Equal(t, "Application", actor["type"])
	assert.Equal(t, "https://example.com/i/actor/inbox", actor["inbox"])
	assert.Equal(t, "https://example.com/i/actor/outbox", actor["outbox"])
	assert.Equal(t, "example.com", actor["preferredUsername"])
	assert.Equal(t, true, actor["manuallyApprovesFollowers"])
	assert.Equal(t, "https://example.com/site/kb/instance-actor", actor["url"])

	publicKey, ok := actor["publicKey"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "https://example.com/i/actor#main-key", publicKey["id"])
	assert.Equal(t, "https://example.com/i/actor", publicKey["owner"])
	assert.Equal(t, "test-public-key", publicKey["publicKeyPem"])

	context, ok := actor["@context"].([]any)
	assert.True(t, ok)
	assert.Contains(t, context, "https://www.w3.org/ns/activitystreams")
	assert.Contains(t, context, "https://w3id.org/security/v1")
}
