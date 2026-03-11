package data

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	App     AppConfig     `mapstructure:"app"`
	Server  ServerConfig  `mapstructure:"server"`
	Service ServiceConfig `mapstructure:"service"`
}

type AppConfig struct {
	Name       string           `mapstructure:"name" env:"APP_NAME"`
	Env        string           `mapstructure:"env" env:"APP_ENV"`
	URL        string           `mapstructure:"url" env:"APP_URL"`
	Federation FederationConfig `mapstructure:"federation"`
}

type ServerConfig struct {
	API APIServerConfig `mapstructure:"api"`
}

type APIServerConfig struct {
	Bind string `mapstrcuture:"bind"`
}

type ServiceConfig struct {
	Database      DatabaseConfig      `mapstructure:"database"`
	OpenTelemetry OpenTelemetryConfig `mapstructure:"open_telemetry"`
}

type DatabaseConfig struct {
	SQL   SQLDBConfig `mapstructure:"sql"`
	Redis RedisConfig `mapstructure:"redis"`
}

type SQLDBConfig struct {
	Driver string `mapstructure:"driver" env:"DB_DRIVER"`
	DSN    string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr" env:"REDIS_HOST"`
	Password string `mapstructure:"password" env:"REDIS_PASSWORD"`
}

type OpenTelemetryConfig struct {
	TracingEnabled  bool   `mapstructure:"tracing_enabled"`
	TracingEndpoint string `mapstructure:"tracing_endpoint"`
}

type FederationConfig struct {
	NodeInfo NodeInfoConfig `mapstructure:"nodeinfo"`
}

type NodeInfoConfig struct {
	Enabled bool `mapstructure:"enabled"`
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
