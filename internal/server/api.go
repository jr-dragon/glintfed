package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/riandyrn/otelchi"
	"glintfed.org/internal/data"
	"glintfed.org/internal/service/healthcheck"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewAPIServer(cfg data.Config) *http.Server {
	mux := chi.NewRouter()

	mux.Use(
		otelchi.Middleware("API", otelchi.WithChiRoutes(mux)),
		middleware.Logger,
		middleware.Recoverer,
	)

	mux.Get("/api/service/health-check", healthcheck.Get)

	return &http.Server{
		Addr:    cfg.Server.API.Bind,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
}
