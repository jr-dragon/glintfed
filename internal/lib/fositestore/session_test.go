package fositestore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalScopes(t *testing.T) {
	assert.Equal(t, "[]", marshalScopes(nil))
	assert.Equal(t, "[]", marshalScopes([]string{}))
	assert.Equal(t, `["read"]`, marshalScopes([]string{"read"}))
	assert.Equal(t, `["read","write"]`, marshalScopes([]string{"read", "write"}))
}

func TestUnmarshalScopes(t *testing.T) {
	assert.Nil(t, unmarshalScopes(""))
	assert.Equal(t, []string{}, unmarshalScopes("[]"))
	assert.Equal(t, []string{"read"}, unmarshalScopes(`["read"]`))
	assert.Equal(t, []string{"read", "write"}, unmarshalScopes(`["read","write"]`))
	assert.Nil(t, unmarshalScopes("invalid-json"))
}

func TestMarshalUnmarshalRoundtrip(t *testing.T) {
	scopes := []string{"read", "write", "follow", "push"}
	assert.Equal(t, scopes, unmarshalScopes(marshalScopes(scopes)))
}
