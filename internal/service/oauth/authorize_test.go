package oauth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSvc_Authorize_ReturnsAccessDenied(t *testing.T) {
	env := newTestEnv(t)

	tests := []struct {
		name  string
		query string
	}{
		{
			name:  "missing required params",
			query: "",
		},
		{
			name:  "valid authorization_code request - login UI not implemented",
			query: "?response_type=code&client_id=1&redirect_uri=https://example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/oauth/authorize"+tt.query, nil)
			w := httptest.NewRecorder()

			env.svc.Authorize(w, req)

			// The authorize endpoint always returns an error (Login UI not yet implemented).
			assert.NotEqual(t, http.StatusOK, w.Code)
		})
	}
}
