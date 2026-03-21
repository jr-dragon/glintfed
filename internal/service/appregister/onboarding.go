package appregister

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
	usermodel "glintfed.org/internal/model/user"
	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

type onboardingRequest struct {
	Email      string `json:"email"       validate:"required,email"`
	VerifyCode string `json:"verify_code" validate:"required,len=6,numeric"`
	Username   string `json:"username"    validate:"required,min=2,max=30,username"`
	Name       string `json:"name"        validate:"omitempty"`
	Password   string `json:"password"    validate:"required"`
}

type onboardingUserResponse struct {
	PID      string `json:"pid"`
	Username string `json:"username"`
}

type onboardingResponse struct {
	Status       string                 `json:"status"`
	TokenType    string                 `json:"token_type"`
	Domain       string                 `json:"domain"`
	ExpiresIn    int64                  `json:"expires_in"`
	AccessToken  string                 `json:"access_token"`
	RefreshToken string                 `json:"refresh_token"`
	ClientID     string                 `json:"client_id"`
	ClientSecret string                 `json:"client_secret"`
	Scope        []string               `json:"scope"`
	User         onboardingUserResponse `json:"user"`
	Account      json.RawMessage        `json:"account,omitempty"`
}

func (s *svc) Onboarding(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "AppRegister.Onboarding")
	defer span.End()

	if !s.cfg.App.Auth.InAppRegistration {
		http.NotFound(w, r)
		return
	}

	if !s.cfg.App.Auth.EnableRegistration {
		http.Error(w, "", http.StatusForbidden)
		return
	}

	var req onboardingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := s.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Dynamic length constraints sourced from config.
	if len(req.Password) < s.cfg.App.MinPasswordLength {
		http.Error(w, fmt.Sprintf("password must be at least %d characters", s.cfg.App.MinPasswordLength), http.StatusBadRequest)
		return
	}
	if s.cfg.App.MaxNameLength > 0 && len(req.Name) > s.cfg.App.MaxNameLength {
		http.Error(w, "name is too long", http.StatusBadRequest)
		return
	}

	email := strings.ToLower(req.Email)

	exists, err := s.arm.VerifyCodeExists(ctx, email, req.VerifyCode)
	if err != nil {
		slog.ErrorContext(ctx, "failed to verify code", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if encErr := json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid verification code, please try again later.",
		}); encErr != nil {
			slog.ErrorContext(ctx, "failed to encode error response", logs.ErrAttr(encErr))
		}
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}

	user, err := s.um.Create(ctx, usermodel.CreateUserParams{
		Name:            req.Name,
		Username:        req.Username,
		Email:           email,
		Password:        req.Password,
		AppRegisterIP:   ip,
		RegisterSource:  "app",
		EmailVerifiedAt: time.Now(),
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to create user", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	scopes := []string{"read", "write", "follow", "push"}
	tokens, err := s.ouc.CreateTokens(ctx, user.ID, scopes)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create OAuth tokens", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := s.arm.DeleteByEmail(ctx, email); err != nil {
		slog.ErrorContext(ctx, "failed to delete app register record", logs.ErrAttr(err))
		// non-fatal: continue and return the successful response
	}

	resp := onboardingResponse{
		Status:       "success",
		TokenType:    "Bearer",
		Domain:       s.cfg.App.Url,
		ExpiresIn:    tokens.ExpiresIn,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ClientID:     tokens.ClientID,
		ClientSecret: tokens.ClientSecret,
		Scope:        scopes,
		User: onboardingUserResponse{
			PID:      fmt.Sprintf("%d", user.ProfileID),
			Username: user.Username,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.ErrorContext(ctx, "failed to encode onboarding response", logs.ErrAttr(err))
	}
}

// validateUsernameTag is a go-playground/validator custom function for the "username" tag.
// It enforces the same rules as PHP's validateUsernameRule:
//   - must not end with .php, .js, or .css
//   - at most one of: dash (-), period (.), underscore (_)
//   - must start and end with an alphanumeric character
//   - after removing special chars, remaining characters must all be alphanumeric
//   - must contain at least one letter
func validateUsernameTag(fl validator.FieldLevel) bool {
	v := fl.Field().String()

	for _, ext := range []string{".php", ".js", ".css"} {
		if strings.HasSuffix(v, ext) {
			return false
		}
	}

	if strings.Count(v, "-")+strings.Count(v, "_")+strings.Count(v, ".") > 1 {
		return false
	}

	first := rune(v[0])
	if !unicode.IsLetter(first) && !unicode.IsDigit(first) {
		return false
	}

	last := rune(v[len(v)-1])
	if !unicode.IsLetter(last) && !unicode.IsDigit(last) {
		return false
	}

	stripped := strings.NewReplacer("_", "", ".", "", "-", "").Replace(v)
	for _, c := range stripped {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}

	for _, c := range v {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false // no letter found
}
