package instanceactor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

func TestProfile(t *testing.T) {
	cfg := &data.Config{
		App: data.AppConfig{
			Url: "https://example.com",
		},
	}

	ia := &ent.InstanceActor{
		PublicKey: "test-public-key",
	}

	iag := &InstanceActorGetterMock{
		GetFunc: func(ctx context.Context) (*ent.InstanceActor, error) {
			return ia, nil
		},
	}

	s := New(cfg, iag)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/i/actor", nil)

	s.Profile(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var res map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, "https://example.com/i/actor", res["id"])
	assert.Equal(t, "Application", res["type"])
	assert.Equal(t, "example.com", res["preferredUsername"])
}

func TestOutbox(t *testing.T) {
	cfg := &data.Config{
		App: data.AppConfig{
			Url: "https://example.com",
		},
	}
	iag := &InstanceActorGetterMock{}
	s := New(cfg, iag)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/i/actor/outbox", nil)

	s.Outbox(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/activity+json", w.Header().Get("Content-Type"))

	var res map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, "https://example.com/i/actor/outbox", res["id"])
	assert.Equal(t, "OrderedCollection", res["type"])
	assert.Equal(t, float64(0), res["totalItems"])
	assert.Equal(t, "https://example.com/i/actor/outbox?page=true", res["first"])
	assert.Equal(t, "https://example.com/i/actor/outbox?min_id=0&page=true", res["last"])

	ctx, ok := res["@context"].([]any)
	assert.True(t, ok)
	assert.Len(t, ctx, 3)
	assert.Equal(t, "https://www.w3.org/ns/activitystreams", ctx[0])
	assert.Equal(t, "https://w3id.org/security/v1", ctx[1])
}
