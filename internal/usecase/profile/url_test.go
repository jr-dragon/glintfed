package profile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"glintfed.org/ent"
)

func TestUsecase_Url(t *testing.T) {
	uc := &Usecase{}

	t.Run("with_remote_url", func(t *testing.T) {
		p := &ent.Profile{
			RemoteURL: "https://example.com/@alice",
			Username:  "alice",
		}
		assert.Equal(t, "https://example.com/@alice", uc.Url(p))
	})

	t.Run("without_remote_url", func(t *testing.T) {
		p := &ent.Profile{
			Username: "alice",
		}
		assert.Equal(t, "alice", uc.Url(p))
	})

	t.Run("with_suffixes", func(t *testing.T) {
		p := &ent.Profile{
			Username: "alice",
		}
		assert.Equal(t, "alice/posts/1", uc.Url(p, "/posts", "/1"))
	})
}

func TestUsecase_Permalink(t *testing.T) {
	uc := &Usecase{}

	t.Run("with_remote_url", func(t *testing.T) {
		p := &ent.Profile{
			RemoteURL: "https://example.com/@alice",
			Username:  "alice",
		}
		assert.Equal(t, "https://example.com/@alice", uc.Permalink(p))
	})

	t.Run("without_remote_url", func(t *testing.T) {
		p := &ent.Profile{
			Username: "alice",
		}
		assert.Equal(t, "users/alice", uc.Permalink(p))
	})

	t.Run("with_suffixes", func(t *testing.T) {
		p := &ent.Profile{
			Username: "alice",
		}
		assert.Equal(t, "users/alice/posts/1", uc.Permalink(p, "/posts", "/1"))
	})
}
