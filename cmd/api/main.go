package main

import (
	"flag"
	"log/slog"

	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/logs"
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
	}

	_ = cfg
}
