package appregister

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	usermodel "glintfed.org/internal/model/user"
	"glintfed.org/internal/usecase/oauth"
)

func TestSvc_Onboarding(t *testing.T) {
	v := validator.New()
	_ = v.RegisterValidation("username", validateUsernameTag)

	defaultCfg := data.AuthConfig{
		InAppRegistration:  true,
		EnableRegistration: true,
	}

	tests := []struct {
		name           string
		authCfg        data.AuthConfig
		appCfg         data.AppConfig
		reqBody        any
		mockVerifyCode func(ctx context.Context, email, code string) (bool, error)
		mockCreateUser func(ctx context.Context, params usermodel.CreateUserParams) (*ent.User, error)
		mockCreateToks func(ctx context.Context, userID uint64, scopes []string) (*oauth.TokenResult, error)
		expectedStatus int
		expectedBody   any
	}{
		{
			name: "InAppRegistration is false",
			authCfg: data.AuthConfig{
				InAppRegistration: false,
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "EnableRegistration is false",
			authCfg: data.AuthConfig{
				EnableRegistration: false,
				InAppRegistration:  true,
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "Invalid JSON body",
			authCfg:        defaultCfg,
			reqBody:        "invalid json",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "Validation fails - invalid username (.php)",
			authCfg: defaultCfg,
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "test.php",
				Password:   "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "Validation fails - too many special chars",
			authCfg: defaultCfg,
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "test-user.name",
				Password:   "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "Validation fails - starts with dash",
			authCfg: defaultCfg,
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "-test",
				Password:   "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "Validation fails - no letter",
			authCfg: defaultCfg,
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "123456",
				Password:   "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "Password too short",
			authCfg: defaultCfg,
			appCfg: data.AppConfig{
				MinPasswordLength: 10,
			},
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "testuser",
				Password:   "short",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "Name too long",
			authCfg: defaultCfg,
			appCfg: data.AppConfig{
				MaxNameLength: 5,
			},
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "testuser",
				Name:       "TooLongName",
				Password:   "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "VerifyCodeExists returns error",
			authCfg: defaultCfg,
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "testuser",
				Password:   "password123",
			},
			mockVerifyCode: func(ctx context.Context, email, code string) (bool, error) {
				return false, errors.New("db error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:    "Invalid verification code",
			authCfg: defaultCfg,
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "654321", // valid format but wrong code
				Username:   "testuser",
				Password:   "password123",
			},
			mockVerifyCode: func(ctx context.Context, email, code string) (bool, error) {
				return false, nil
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{"status": "error", "message": "Invalid verification code, please try again later."},
		},
		{
			name:    "User creation fails",
			authCfg: defaultCfg,
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "testuser",
				Password:   "password123",
			},
			mockVerifyCode: func(ctx context.Context, email, code string) (bool, error) {
				return true, nil
			},
			mockCreateUser: func(ctx context.Context, params usermodel.CreateUserParams) (*ent.User, error) {
				return nil, errors.New("create error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:    "Token creation fails",
			authCfg: defaultCfg,
			reqBody: onboardingRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
				Username:   "testuser",
				Password:   "password123",
			},
			mockVerifyCode: func(ctx context.Context, email, code string) (bool, error) {
				return true, nil
			},
			mockCreateUser: func(ctx context.Context, params usermodel.CreateUserParams) (*ent.User, error) {
				return &ent.User{ID: 123}, nil
			},
			mockCreateToks: func(ctx context.Context, userID uint64, scopes []string) (*oauth.TokenResult, error) {
				return nil, errors.New("token error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:    "Successful onboarding",
			authCfg: defaultCfg,
			appCfg: data.AppConfig{
				Url: "https://glintfed.org",
			},
			reqBody: onboardingRequest{
				Email:      "Test@Example.Com",
				VerifyCode: "123456",
				Username:   "test_user",
				Name:       "Test User",
				Password:   "password123",
			},
			mockVerifyCode: func(ctx context.Context, email, code string) (bool, error) {
				return true, nil
			},
			mockCreateUser: func(ctx context.Context, params usermodel.CreateUserParams) (*ent.User, error) {
				assert.Equal(t, "test@example.com", params.Email)
				return &ent.User{ID: 123, ProfileID: 456, Username: "test_user"}, nil
			},
			mockCreateToks: func(ctx context.Context, userID uint64, scopes []string) (*oauth.TokenResult, error) {
				assert.Equal(t, uint64(123), userID)
				return &oauth.TokenResult{
					AccessToken:  "at",
					RefreshToken: "rt",
					ExpiresIn:    3600,
					ClientID:     "cid",
					ClientSecret: "cs",
				}, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody: onboardingResponse{
				Status:       "success",
				TokenType:    "Bearer",
				Domain:       "https://glintfed.org",
				ExpiresIn:    3600,
				AccessToken:  "at",
				RefreshToken: "rt",
				ClientID:     "cid",
				ClientSecret: "cs",
				Scope:        []string{"read", "write", "follow", "push"},
				User: onboardingUserResponse{
					PID:      "456",
					Username: "test_user",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arm := &AppRegisterModelMock{
				VerifyCodeExistsFunc: tt.mockVerifyCode,
				DeleteByEmailFunc: func(ctx context.Context, email string) error {
					return nil
				},
			}
			um := &UserModelMock{
				CreateFunc: tt.mockCreateUser,
			}
			ouc := &OAuthUsecaseMock{
				CreateTokensFunc: tt.mockCreateToks,
			}

			cfg := &data.Config{
				App: tt.appCfg,
			}
			cfg.App.Auth = tt.authCfg

			s := &svc{
				cfg:      cfg,
				validate: v,
				arm:      arm,
				um:       um,
				ouc:      ouc,
			}

			var body []byte
			if str, ok := tt.reqBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.reqBody)
			}

			req := httptest.NewRequest(http.MethodPost, "/onboarding", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			s.Onboarding(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedBody != nil {
				if expectedMap, ok := tt.expectedBody.(map[string]string); ok {
					var actualMap map[string]string
					err := json.Unmarshal(w.Body.Bytes(), &actualMap)
					assert.NoError(t, err)
					assert.Equal(t, expectedMap, actualMap)
				} else {
					var actualResp onboardingResponse
					err := json.Unmarshal(w.Body.Bytes(), &actualResp)
					assert.NoError(t, err)
					assert.Equal(t, tt.expectedBody, actualResp)
				}
			}
		})
	}
}

func TestValidateUsernameTag(t *testing.T) {
	v := validator.New()
	_ = v.RegisterValidation("username", validateUsernameTag)

	type testStruct struct {
		Username string `validate:"username"`
	}

	tests := []struct {
		username string
		valid    bool
	}{
		{"testuser", true},
		{"test_user", true},
		{"test-user", true},
		{"test.user", true},
		{"test123", true},
		{"user123", true},
		{"a1", true},
		{"1a", true},
		{"test.php", false},
		{"test.js", false},
		{"test.css", false},
		{"test--user", false},
		{"test__user", false},
		{"test..user", false},
		{"_test", false},
		{"test_", false},
		{"-test", false},
		{"test-", false},
		{".test", false},
		{"test.", false},
		{"12345", false}, // no letter
		{"abc!", false},  // invalid char
		{"a.b", true},
		{"a_b", true},
		{"a-b", true},
		{"a.b_c", false}, // too many special chars
	}

	for _, tt := range tests {
		err := v.Struct(testStruct{Username: tt.username})
		if tt.valid {
			assert.NoError(t, err, "expected %s to be valid", tt.username)
		} else {
			assert.Error(t, err, "expected %s to be invalid", tt.username)
		}
	}
}
