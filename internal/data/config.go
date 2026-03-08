package data

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	API APIServerConfig `mapstructure:"api"`
}

type APIServerConfig struct {
	Bind string `mapstrcuture:"bind"`
}

func NewConfig(paths ...string) (cfg Config, err error) {
	config.WithOptions(config.ParseEnv, config.ParseTime)
	config.AddDriver(yaml.Driver)

	if err = config.LoadFiles(paths...); err != nil {
		return
	}
	if err = config.BindStruct("", &cfg); err != nil {
		return
	}

	return
}
