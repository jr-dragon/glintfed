package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"

	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/server"
)

// Name is the name of the application.
var Name string

// Version is the version of the application.
var Version string

var (
	flagCfgPath string
)

func init() {
	flag.StringVar(&flagCfgPath, "config", "config.yaml", "config file path")
}

func main() {
	flag.Parse()

	cfg, err := data.NewConfig(flagCfgPath)
	if err != nil {
		slog.Error("failed to load config", logs.ErrAttr(err))
		return
	}

	otelCleanup, err := setupOTelSDK(context.Background(), cfg)
	if err != nil {
		slog.Error("failed to init open telemetry sdk", logs.ErrAttr(err))
		return
	}
	defer func() {
		if err := errors.Join(err, otelCleanup(context.Background())); err != nil {
			slog.Error("failed to cleanup open telemetry sdk", logs.ErrAttr(err))
		}
	}()

	client, clientCleanup, err := data.NewClient(cfg)
	if err != nil {
		slog.Error("failed to init clients", logs.ErrAttr(err))
		return
	}
	defer clientCleanup()

	if err := InitApp(cfg, client).ListenAndServe(); err != nil {
		slog.Error("failed to start api server", logs.ErrAttr(err))
		return
	}
}

type app struct {
	*http.Server
}

func newapp(cfg *data.Config, svcs *server.Services) *app {
	return &app{
		Server: server.NewAPIServer(cfg, svcs),
	}
}
