package fositestore

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ory/fosite"

	"glintfed.org/ent"
	"glintfed.org/ent/oauthclient"
)

// FositeClient wraps *ent.OauthClient and implements fosite.Client.
type FositeClient struct {
	*ent.OauthClient
}

// GetID returns the client ID as a string.
func (c *FositeClient) GetID() string {
	return strconv.FormatUint(c.ID, 10)
}

// GetHashedSecret returns the hashed client secret.
func (c *FositeClient) GetHashedSecret() []byte {
	return []byte(c.Secret)
}

// GetRedirectURIs returns redirect URIs split from the newline-delimited redirect field.
func (c *FositeClient) GetRedirectURIs() []string {
	if c.Redirect == "" {
		return nil
	}
	return strings.Split(c.Redirect, "\n")
}

// GetGrantTypes returns the allowed grant types.
func (c *FositeClient) GetGrantTypes() fosite.Arguments {
	return fosite.Arguments(c.GrantTypes)
}

// GetResponseTypes returns the allowed response types.
func (c *FositeClient) GetResponseTypes() fosite.Arguments {
	return fosite.Arguments(c.ResponseTypes)
}

// GetScopes returns the allowed scopes.
func (c *FositeClient) GetScopes() fosite.Arguments {
	return fosite.Arguments(c.Scopes)
}

// IsPublic returns whether the client is public (no client secret required).
func (c *FositeClient) IsPublic() bool {
	return c.Public
}

// GetAudience returns the allowed audience.
func (c *FositeClient) GetAudience() fosite.Arguments {
	return fosite.Arguments(c.Audience)
}

// GetClient retrieves a client by its string ID.
//
//	SELECT * FROM oauth_clients WHERE id = ? LIMIT 1
func (s *Store) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	idNum, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid client id %q", fosite.ErrNotFound, id)
	}
	c, err := s.db.OauthClient.Get(ctx, idNum)
	if ent.IsNotFound(err) {
		return nil, fosite.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if c.Revoked {
		return nil, fosite.ErrInvalidClient
	}
	return &FositeClient{c}, nil
}

// GetClientByNumericID retrieves a client by its numeric ID.
//
//	SELECT * FROM oauth_clients WHERE id = ? LIMIT 1
func (s *Store) GetClientByNumericID(ctx context.Context, id uint64) (fosite.Client, error) {
	c, err := s.db.OauthClient.Get(ctx, id)
	if ent.IsNotFound(err) {
		return nil, fosite.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if c.Revoked {
		return nil, fosite.ErrInvalidClient
	}
	return &FositeClient{c}, nil
}

// GetPersonalAccessClient retrieves the first personal access client record.
//
//	SELECT * FROM oauth_clients WHERE personal_access_client = true LIMIT 1
func (s *Store) GetPersonalAccessClient(ctx context.Context) (*FositeClient, error) {
	c, err := s.db.OauthClient.Query().
		Where(oauthclient.PersonalAccessClient(true)).
		First(ctx)
	if ent.IsNotFound(err) {
		return nil, fosite.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &FositeClient{c}, nil
}
