package ent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const appurl = "https://glitfed.test"

func TestProfile_Url(t *testing.T) {
	t.Run("with_remote_url", func(t *testing.T) {
		p := &Profile{
			RemoteURL: "https://example.com/@alice",
			Username:  "alice",
		}
		assert.Equal(t, "https://example.com/@alice", p.Url(appurl))
	})

	t.Run("without_remote_url", func(t *testing.T) {
		p := &Profile{
			Username: "alice",
		}
		assert.Equal(t, appurl+"/alice", p.Url(appurl))
	})

	t.Run("with_suffixes", func(t *testing.T) {
		p := &Profile{
			Username: "alice",
		}
		assert.Equal(t, appurl+"/alice/posts/1.png", p.Url(appurl, "/posts", "/1", ".png"))
	})
}

func TestProfile_Permalink(t *testing.T) {
	t.Run("with_remote_url", func(t *testing.T) {
		p := &Profile{
			RemoteURL: "https://example.com/@alice",
			Username:  "alice",
		}
		assert.Equal(t, "https://example.com/@alice", p.Permalink(appurl))
	})

	t.Run("without_remote_url", func(t *testing.T) {
		p := &Profile{
			Username: "alice",
		}
		assert.Equal(t, appurl+"/users/alice", p.Permalink(appurl))
	})

	t.Run("with_suffixes", func(t *testing.T) {
		p := &Profile{
			Username: "alice",
		}
		assert.Equal(t, appurl+"/users/alice/posts/1.png", p.Permalink(appurl, "/posts", "/1", ".png"))
	})
}
