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
	"glintfed.org/internal/data"
)

func TestSvc_VerifyCode(t *testing.T) {
	v := validator.New()

	tests := []struct {
		name           string
		cfg            data.AuthConfig
		reqBody        any
		mockVerifyCode func(ctx context.Context, email, code string) (bool, error)
		expectedStatus int
		expectedBody   any
	}{
		{
			name: "InAppRegistration is false",
			cfg: data.AuthConfig{
				InAppRegistration:  false,
				EnableRegistration: true,
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "EnableRegistration is false",
			cfg: data.AuthConfig{
				InAppRegistration:  true,
				EnableRegistration: false,
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "Invalid JSON body",
			cfg: data.AuthConfig{
				InAppRegistration:  true,
				EnableRegistration: true,
			},
			reqBody:        "invalid json",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Validation fails - missing email",
			cfg: data.AuthConfig{
				InAppRegistration:  true,
				EnableRegistration: true,
			},
			reqBody: verifyCodeRequest{
				VerifyCode: "123456",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Validation fails - invalid verify code",
			cfg: data.AuthConfig{
				InAppRegistration:  true,
				EnableRegistration: true,
			},
			reqBody: verifyCodeRequest{
				Email:      "test@example.com",
				VerifyCode: "123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "VerifyCodeExists returns error",
			cfg: data.AuthConfig{
				InAppRegistration:  true,
				EnableRegistration: true,
			},
			reqBody: verifyCodeRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
			},
			mockVerifyCode: func(ctx context.Context, email, code string) (bool, error) {
				return false, errors.New("db error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "VerifyCodeExists returns false (not found or expired)",
			cfg: data.AuthConfig{
				InAppRegistration:  true,
				EnableRegistration: true,
			},
			reqBody: verifyCodeRequest{
				Email:      "test@example.com",
				VerifyCode: "123456",
			},
			mockVerifyCode: func(ctx context.Context, email, code string) (bool, error) {
				return false, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   verifyCodeResponse{Status: "error"},
		},
		{
			name: "VerifyCodeExists returns true (success)",
			cfg: data.AuthConfig{
				InAppRegistration:  true,
				EnableRegistration: true,
			},
			reqBody: verifyCodeRequest{
				Email:      "Test@Example.Com", // Mixed case to test normalization
				VerifyCode: "123456",
			},
			mockVerifyCode: func(ctx context.Context, email, code string) (bool, error) {
				if email == "test@example.com" && code == "123456" {
					return true, nil
				}
				return false, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   verifyCodeResponse{Status: "success"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arm := &AppRegisterModelMock{
				VerifyCodeExistsFunc: tt.mockVerifyCode,
			}
			s := &svc{
				cfg: &data.Config{
					App: data.AppConfig{
						Auth: tt.cfg,
					},
				},
				validate: v,
				arm:      arm,
			}

			var body []byte
			if s, ok := tt.reqBody.(string); ok {
				body = []byte(s)
			} else {
				body, _ = json.Marshal(tt.reqBody)
			}

			req := httptest.NewRequest(http.MethodPost, "/verify-code", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			s.VerifyCode(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedBody != nil {
				var resp verifyCodeResponse
				err := json.Unmarshal(w.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, resp)
			}
		})
	}
}
