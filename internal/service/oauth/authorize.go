package oauth

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/ory/fosite"

	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

// Authorize handles GET /oauth/authorize.
//
// It validates the authorize request via fosite, then redirects the user to
// the configured LoginUrl so the frontend can authenticate them. The full
// original query string is forwarded as a `next` parameter so the frontend
// can redirect back to this endpoint after login.
//
// If LoginUrl is not configured, the endpoint returns ErrAccessDenied.
func (s *svc) Authorize(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "OAuth.Authorize")
	defer span.End()

	authReq, err := s.provider.NewAuthorizeRequest(ctx, r)
	if err != nil {
		s.provider.WriteAuthorizeError(ctx, w, authReq, err)
		return
	}

	if s.loginURL == "" {
		slog.ErrorContext(ctx, "OAuth.Authorize: login_url not configured")
		s.provider.WriteAuthorizeError(ctx, w, authReq, fosite.ErrServerError.WithDescription("login URL not configured"))
		return
	}

	// Build the return URL: the original /oauth/authorize request with all its params.
	returnURL := s.appURL + "/oauth/authorize?" + r.URL.RawQuery

	loginRedirect, err := url.Parse(s.loginURL)
	if err != nil {
		slog.ErrorContext(ctx, "OAuth.Authorize: invalid login_url", logs.ErrAttr(err))
		s.provider.WriteAuthorizeError(ctx, w, authReq, fosite.ErrServerError.WithDescription("invalid login URL"))
		return
	}
	q := loginRedirect.Query()
	q.Set("next", returnURL)
	loginRedirect.RawQuery = q.Encode()

	http.Redirect(w, r, loginRedirect.String(), http.StatusSeeOther)
}
