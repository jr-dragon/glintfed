package oauth

import (
	"net/http"

	"glintfed.org/internal/service/internal"
)

// Revoke handles POST /oauth/revoke.
func (s *svc) Revoke(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "OAuth.Revoke")
	defer span.End()

	err := s.provider.NewRevocationRequest(ctx, r)
	s.provider.WriteRevocationResponse(ctx, w, err)
}
