package oauth

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/ory/fosite"

	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

const oobRedirectURI = "urn:ietf:wg:oauth:2.0:oob"

// Authorize handles GET /oauth/authorize.
func (s *svc) Authorize(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "OAuth.Authorize")
	defer span.End()

	authReq, err := s.provider.NewAuthorizeRequest(ctx, r)
	if err != nil {
		s.provider.WriteAuthorizeError(ctx, w, authReq, err)
		return
	}

	// TODO: Add actual user authentication/session check here.
	// For now, we return an error indicating the user must authenticate.
	// A full implementation would redirect to a login page and resume the flow.
	s.provider.WriteAuthorizeError(ctx, w, authReq, fosite.ErrAccessDenied.WithDescription("user authentication not implemented"))
	_ = logs.ErrAttr
	_ = slog.ErrorContext
	_ = url.QueryEscape
	_ = oobRedirectURI
}
