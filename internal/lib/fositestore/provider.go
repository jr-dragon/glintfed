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
	tokenDays := cfg.App.Auth.OAuth.AccessTokenLifespanDays
	refreshDays := cfg.App.Auth.OAuth.RefreshTokenLifespanDays

	if tokenDays <= 0 {
		tokenDays = 365
	}
	if refreshDays <= 0 {
		refreshDays = 400
	}

	fositeCfg := &fosite.Config{
		AccessTokenLifespan:        time.Duration(tokenDays) * 24 * time.Hour,
		RefreshTokenLifespan:       time.Duration(refreshDays) * 24 * time.Hour,
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
