package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ory/fosite"
)

type ctxKey int

const (
	// CtxKeySubject is the context key for the authenticated user subject (user ID as string).
	CtxKeySubject ctxKey = iota
	// CtxKeyScopes is the context key for the granted OAuth scopes.
	CtxKeyScopes
)

// OAuth2Auth returns a middleware that validates Bearer tokens using the given fosite provider.
// On success, it injects the subject and granted scopes into the request context.
func OAuth2Auth(provider fosite.OAuth2Provider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractBearerToken(r)
			if token == "" {
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, `{"error":"missing_token"}`, http.StatusUnauthorized)
				return
			}

			_, ar, err := provider.IntrospectToken(
				r.Context(), token, fosite.AccessToken, &fosite.DefaultSession{},
			)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), CtxKeySubject, ar.GetSession().GetSubject())
			ctx = context.WithValue(ctx, CtxKeyScopes, ar.GetGrantedScopes())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractBearerToken(r *http.Request) string {
	token, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !ok {
		return ""
	}
	return token
}

// SubjectFromContext retrieves the authenticated subject from the context.
func SubjectFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(CtxKeySubject).(string); ok {
		return v
	}
	return ""
}

// ScopesFromContext retrieves the granted scopes from the context.
func ScopesFromContext(ctx context.Context) []string {
	if v, ok := ctx.Value(CtxKeyScopes).([]string); ok {
		return v
	}
	return nil
}
