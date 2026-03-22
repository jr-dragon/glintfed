package fositestore

import (
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"

	"glintfed.org/internal/data"
)

// NewOAuth2Provider builds a fosite.OAuth2Provider using the given Store and config.
// The Store already holds the HMAC strategy derived from the config secret.
func NewOAuth2Provider(store *Store, cfg *data.Config) fosite.OAuth2Provider {
	tokenTTL := cfg.App.Auth.OAuth.AccessTokenLifespan
	refreshTTL := cfg.App.Auth.OAuth.RefreshTokenLifespan

	if tokenTTL <= 0 {
		tokenTTL = 365 * 24 * time.Hour
	}
	if refreshTTL <= 0 {
		refreshTTL = 400 * 24 * time.Hour
	}

	fositeCfg := &fosite.Config{
		AccessTokenLifespan:        tokenTTL,
		RefreshTokenLifespan:       refreshTTL,
		AuthorizeCodeLifespan:      10 * time.Minute,
		SendDebugMessagesToClients: false,
	}

	strategy := &compose.CommonStrategy{
		CoreStrategy: store.Strategy(),
	}

	return compose.Compose(
		fositeCfg,
		store,
		strategy,
		compose.OAuth2AuthorizeExplicitFactory,
		compose.OAuth2ClientCredentialsGrantFactory,
		compose.OAuth2RefreshTokenGrantFactory,
		compose.OAuth2TokenRevocationFactory,
		compose.OAuth2TokenIntrospectionFactory,
		compose.OAuth2PKCEFactory,
	)
}
