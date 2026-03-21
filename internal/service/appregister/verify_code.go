package appregister

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

type verifyCodeRequest struct {
	Email      string `json:"email"       validate:"required,email"`
	VerifyCode string `json:"verify_code" validate:"required,len=6,numeric"`
}

type verifyCodeResponse struct {
	Status string `json:"status"`
}

func (s *svc) VerifyCode(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "AppRegister.VerifyCode")
	defer span.End()

	if !s.cfg.App.Auth.InAppRegistration {
		http.NotFound(w, r)
		return
	}

	if !s.cfg.App.Auth.EnableRegistration {
		http.Error(w, "", http.StatusForbidden)
		return
	}

	var req verifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := s.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := strings.ToLower(req.Email)

	exists, err := s.arm.VerifyCodeExists(ctx, email, req.VerifyCode)
	if err != nil {
		slog.ErrorContext(ctx, "failed to verify code", logs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	status := "error"
	if exists {
		status = "success"
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(verifyCodeResponse{Status: status}); err != nil {
		slog.ErrorContext(ctx, "failed to encode verify code response", logs.ErrAttr(err))
	}
}
