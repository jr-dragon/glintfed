package appregister

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	usermodel "glintfed.org/internal/model/user"
)

type Service interface {
	VerifyCode(w http.ResponseWriter, r *http.Request)
	Onboarding(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_app_register_model.go . AppRegisterModel
type AppRegisterModel interface {
	// VerifyCodeExists checks whether an AppRegister record exists for the given email and
	// verify_code that was created within the past 90 days.
	VerifyCodeExists(ctx context.Context, email, code string) (bool, error)
	// DeleteByEmail removes the AppRegister record for the given email after onboarding completes.
	DeleteByEmail(ctx context.Context, email string) error
}

//go:generate go tool moq -rm -out mock_user_model.go . UserModel
type UserModel interface {
	// Create inserts a new user with the given parameters. The implementation is responsible
	// for hashing the plaintext password before storing it.
	Create(ctx context.Context, params usermodel.CreateUserParams) (*ent.User, error)
}

//go:generate go tool moq -rm -out mock_oauth_model.go . OAuthModel
type OAuthModel interface {
	// CreateTokens creates an OAuth access token and a refresh token for the given user ID
	// with the specified scopes, and returns the resulting token details.
	CreateTokens(ctx context.Context, userID uint64, scopes []string) (*OAuthTokenResult, error)
}

//go:generate go tool moq -rm -out mock_account_getter.go . AccountGetter
type AccountGetter interface {
	// GetByProfileID returns the Mastodon-compatible account representation for the given
	// profile ID as a raw JSON value.
	GetByProfileID(ctx context.Context, profileID uint64) (json.RawMessage, error)
}

// OAuthTokenResult contains the OAuth token details returned after successful onboarding.
type OAuthTokenResult struct {
	AccessToken  string
	RefreshToken string
	ClientID     string
	ClientSecret string
	ExpiresIn    int64
}

func New(cfg *data.Config, arm AppRegisterModel, um UserModel, om OAuthModel, ag AccountGetter) Service {
	v := validator.New()
	// Register custom tag "username" implementing PHP's validateUsernameRule logic.
	_ = v.RegisterValidation("username", validateUsernameTag)

	return &svc{
		cfg:      cfg,
		validate: v,
		arm:      arm,
		um:       um,
		om:       om,
		ag:       ag,
	}
}

type svc struct {
	cfg      *data.Config
	validate *validator.Validate
	arm      AppRegisterModel
	um       UserModel
	om       OAuthModel
	ag       AccountGetter
}
