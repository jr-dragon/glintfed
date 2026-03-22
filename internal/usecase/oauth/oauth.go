package oauth

import (
	"context"
	"crypto/rand"
	"fmt"
	"strconv"
	"time"

	"github.com/ory/fosite"

	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/fositestore"
)

// Usecase handles direct token issuance for users (bypassing standard OAuth2 flows).
type Usecase struct {
	store    *fositestore.Store
	clientID string
	tokenTTL time.Duration
}

// NewUsecase creates a new Usecase using configuration for the personal client ID and token TTL.
func NewUsecase(store *fositestore.Store, cfg *data.Config) *Usecase {
	ttl := cfg.App.Auth.OAuth.AccessTokenLifespan
	if ttl <= 0 {
		ttl = 365 * 24 * time.Hour
	}
	return &Usecase{
		store:    store,
		clientID: cfg.App.Auth.OAuth.PersonalClientID,
		tokenTTL: ttl,
	}
}

// TokenResult contains the OAuth token details issued after successful authentication.
type TokenResult struct {
	AccessToken  string
	RefreshToken string
	ClientID     string
	ClientSecret string
	ExpiresIn    int64
}

// CreateTokens directly issues an access token and refresh token for the given user ID
// without going through the standard OAuth2 authorization flow.
// This corresponds to pixelfed's $user->createToken('Pixelfed App', scopes).
func (uc *Usecase) CreateTokens(ctx context.Context, userID uint64, scopes []string) (*TokenResult, error) {
	subject := strconv.FormatUint(userID, 10)

	client, err := uc.store.GetClient(ctx, uc.clientID)
	if err != nil {
		return nil, fmt.Errorf("get personal access client: %w", err)
	}

	now := time.Now()
	requestID := generateRequestID()

	session := &fosite.DefaultSession{
		Subject:  subject,
		Username: subject,
		ExpiresAt: map[fosite.TokenType]time.Time{
			fosite.AccessToken:  now.Add(uc.tokenTTL),
			fosite.RefreshToken: now.Add(uc.tokenTTL + 35*24*time.Hour),
		},
	}

	req := fosite.NewRequest()
	req.ID = requestID
	req.Client = client
	req.RequestedAt = now
	req.Session = session
	req.SetRequestedScopes(fosite.Arguments(scopes))
	for _, s := range scopes {
		req.GrantScope(s)
	}

	accessToken, refreshToken, err := uc.store.CreatePersonalAccessTokens(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create tokens: %w", err)
	}

	fositeClient, ok := client.(*fositestore.FositeClient)
	if !ok {
		return nil, fmt.Errorf("unexpected client type")
	}

	return &TokenResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ClientID:     uc.clientID,
		ClientSecret: string(fositeClient.GetHashedSecret()),
		ExpiresIn:    int64(uc.tokenTTL.Seconds()),
	}, nil
}

func generateRequestID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
