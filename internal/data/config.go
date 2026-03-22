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
	Name        string `mapstructure:"name" env:"APP_NAME"`
	Version     string `mapstructure:"version" env:"APP_VERSION"`
	Env         string `mapstructure:"env" env:"APP_ENV"`
	Url         string `mapstructure:"url" env:"APP_URL"`
	Description string `mapstructure:"description" env:"PF_DESCRIPTION"`

	MediaTypes          string `mapstrucure:"media_types" env:"MEDIA_TYPES"`
	MaxPhotoSize        int    `json:"max_photo_size" env:"MAX_PHOTO_SIZE"`
	MaxCaptionLength    int    `json:"max_caption_length" env:"MAX_CAPTION_LENGTH"`
	MaxAltextLength     int    `json:"max_altext_length" env:"PF_MEDIA_MAX_ALTTEXT_LENGTH"`
	MaxAlbumLength      int    `json:"max_album_length" env:"MAX_ALBUM_LENGTH"`
	ImageQuality        int    `json:"image_quality" env:"IMAGE_QUALITY"`
	MaxCollectionLength int    `json:"max_collection_length" env:"PF_MAX_COLLECTION_LENGTH"`
	OptimizeImage       bool   `json:"optimize_image" env:"PF_OPTIMIZE_IMAGES"`
	OptimizeVideo       bool   `json:"optimize_video" env:"PF_OPTIMIZE_VIDEOS"`
	EnforceAcountLimit  bool   `json:"enforce_account_limit" env:"LIMIT_ACCOUNT_SIZE"`
	CloudStorage        bool   `json:"cloud_storage" env:"PF_ENABLE_CLOUD"`

	MaxAvatarSize     int `json:"max_avatar_size" env:"MAX_AVATAR_SIZE"`
	MaxBioLength      int `json:"max_bio_length" env:"MAX_BIO_LENGTH"`
	MaxNameLength     int `json:"max_name_length" env:"MAX_NAME_LENGTH"`
	MinPasswordLength int `json:"min_password_length" env:"MIN_PASSWORD_LENGTH"`
	MaxAccountSize    int `json:"max_account_size" env:"MAX_ACCOUNT_SIZE"`

	Auth       AuthConfig       `mapstructure:"auth"`
	Instance   InstanceConfig   `mapstructure:"instance"`
	Federation FederationConfig `mapstructure:"federation"`
	Import     ImportConfig     `mapstructure:"import"`
	Media      MediaConfig      `mapstructure:"media"`
	Groups     GroupsConfig     `mapstructure:"groups"`
}

type AuthConfig struct {
	EnableRegistration bool        `mapstructure:"enable_registration" env:"OPEN_REGISTRATION"`
	EnableOAuth        bool        `mapstructure:"enable_oauth" env:"OAUTH_ENABLED"`
	InAppRegistration  bool        `mapstructure:"in_app_registration" env:"APP_REGISTER"`
	OAuth              OAuthConfig `mapstructure:"oauth"`
}

// OAuthConfig holds configuration for the embedded OAuth2 server.
type OAuthConfig struct {
	HMACSecret               string `mapstructure:"hmac_secret" env:"OAUTH_HMAC_SECRET"`
	PersonalClientID         string `mapstructure:"personal_client_id" env:"OAUTH_PERSONAL_CLIENT_ID"`
	AccessTokenLifespanDays  int    `mapstructure:"access_token_lifespan_days" env:"OAUTH_TOKEN_EXPIRATION"`
	RefreshTokenLifespanDays int    `mapstructure:"refresh_token_lifespan_days" env:"OAUTH_REFRESH_EXPIRATION"`
}

type UploaderConfig struct {
}

type ActivitypubConfig struct {
	Enabled      bool `json:"enabled" env:"ACTIVITY_PUB"`
	RemoteFollow bool `json:"remote_follow" env:"AP_REMOTE_FOLLOW"`
	SharedInbox  bool `json:"shared_inbox" env:"AP_SHAREDINBOX"`
	Inbox        bool `json:"inbox" env:"AP_INBOX"`
}

type InstanceConfig struct {
	HasLegalNotice bool           `mapstructure:"has_legal_notice" env:"INSTANCE_LEGAL_NOTICE"`
	Username       UsernameConfig `mapstructure:"username"`
	Stories        StoriesConfig  `mapstructure:"stories"`
	Label          LabelConfig    `mapstructure:"label"`
}

type UsernameConfig struct {
	Remote RemoteConfig `mapstructure:"remote"`
}

type StoriesConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

type RemoteConfig struct {
	Formats []string `json:"formats"`
	Format  string   `json:"format"`
	Custom  string   `json:"custom" env:"USERNAME_REMOTE_CUSTOM_TEXT"`
}

type ImportConfig struct {
	Instagram InstagramConfig `mapstructure:"instagram"`
}

type InstagramConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

type MediaConfig struct {
	HLS HLSConfig `mapstructure:"hls"`
}

type HLSConfig struct {
	Enabled  bool   `mapstructure:"enabled" env:"MEDIA_HLS_ENABLED"`
	Debug    bool   `mapstructure:"debug" env:"MEDIA_HLS_DEBUG"`
	P2P      bool   `mapstructure:"p2p" env:"MEDIA_HLS_P2P"`
	P2PDebug bool   `mapstructure:"p2p_debug" env:"MEDIA_HLS_P2P_DEBUG"`
	Tracker  string `mapstructure:"tracker" env:"MEDIA_HLS_P2P_TRACKER"`
	Ice      string `mapstructure:"ice" env:"MEDIA_HLS_P2P_ICE_SERVER"`
}

type LabelConfig struct {
	Covid LabelContentConfig `mapstructure:"covid"`
}

type LabelContentConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Org     string `mapstructure:"org"`
	Url     string `mapstructure:"url"`
}

type GroupsConfig struct {
	Enabled bool `mapstructure:"enabled" env:"GROUPS_ENABLED"`
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
	NodeInfo        NodeInfoConfig    `mapstructure:"nodeinfo"`
	Webfinger       WebfingerConfig   `mapstructure:"webfinger"`
	NetworkTimeline bool              `mapstructure:"network_timeline" env:"PF_NETWORK_TIMELINE"`
	Activitypub     ActivitypubConfig `mapstructure:"activitypub"`
}

type NodeInfoConfig struct {
	Enabled bool `mapstructure:"enabled" env:"NODEINFO"`
}

type WebfingerConfig struct {
	Enabled bool `mapstructure:"enabled" env:"WEBFINGER"`
}

func NewConfig(paths ...string) (cfg *Config, err error) {
	cfg = &Config{}
	config.WithOptions(config.ParseEnv, config.ParseTime)
	config.AddDriver(yaml.Driver)

	if err = config.LoadFiles(paths...); err != nil {
		return
	}
	if err = config.BindStruct("", cfg); err != nil {
		return
	}

	return
}
