package config

import (
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type (
	SslMode string
)

//goland:noinspection GoUnusedConst
const (
	SslModeDisable SslMode = "disable"
	SslModeAllow   SslMode = "allow"
	SslModePrefer  SslMode = "prefer"
	SslModeRequire SslMode = "require"
)

var (
	config *Config
	Path   string
	once   sync.Once
)

// GetConfig returns the configuration.
func GetConfig() *Config {
	once.Do(func() {
		if Path == "" {
			Path = "config.toml"
		}
		config = &Config{}

		viper.SetConfigFile(Path)
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("Error unmarshaling config: %v", err)
		}

		log.Println("Loaded config file:", viper.ConfigFileUsed())
	})

	return config
}

// Config represents the top-level configuration structure.
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Security SecurityConfig `mapstructure:"security"`
	Session  SessionConfig  `mapstructure:"session"`
}

// ServerConfig represents the server configuration.
type ServerConfig struct {
	Port  uint16 `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

// DatabaseConfig represents the database configuration.
type DatabaseConfig struct {
	Host     string  `mapstructure:"host"`
	Port     uint16  `mapstructure:"port"`
	User     string  `mapstructure:"user"`
	Password string  `mapstructure:"password"`
	Db       string  `mapstructure:"db"`
	SslMode  SslMode `mapstructure:"ssl_mode"`
}

// SecurityConfig represents the security configuration.
type SecurityConfig struct {
	AllowedHosts       []string `mapstructure:"allowed_hosts"`
	HSTSMaxAge         int      `mapstructure:"hsts_max_age"`
	XSSProtection      string   `mapstructure:"xss_protection"`
	ContentTypeNosniff string   `mapstructure:"content_type_nosniff"`
	XFrameOptions      string   `mapstructure:"x_frame_options"`
	ReferrerPolicy     string   `mapstructure:"referrer_policy"`
	CSPDefaultSrc      string   `mapstructure:"csp_default_src"`
	CSPScriptSrc       string   `mapstructure:"csp_script_src"`
	CSPStyleSrc        string   `mapstructure:"csp_style_src"`
	CSPImgSrc          string   `mapstructure:"csp_img_src"`
	CSPFontSrc         string   `mapstructure:"csp_font_src"`
}

// SessionConfig represents the session configuration.
type SessionConfig struct {
	Key      string        `mapstructure:"key"`
	EncKey   string        `mapstructure:"enc_key"`
	Path     string        `mapstructure:"path"`
	Domain   string        `mapstructure:"domain"`
	MaxAge   int           `mapstructure:"max_age"`
	Secure   bool          `mapstructure:"secure"`
	HttpOnly bool          `mapstructure:"http_only"`
	SameSite http.SameSite `mapstructure:"same_site"`
}

// URL returns the database URL.
func (c DatabaseConfig) URL() *url.URL {
	query := url.Values{}
	query.Set("sslmode", string(c.SslMode))
	return &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.User, c.Password),
		Host:     net.JoinHostPort(c.Host, fmt.Sprintf("%d", c.Port)),
		Path:     c.Db,
		RawQuery: query.Encode(),
	}
}

func (c SecurityConfig) CSP() string {
	return fmt.Sprintf(
		"default-src %s; script-src %s; style-src %s; img-src %s; font-src %s",
		c.CSPDefaultSrc, c.CSPScriptSrc, c.CSPStyleSrc, c.CSPImgSrc, c.CSPFontSrc,
	)
}

// KeyBytes returns the session key as a byte array.
// The key is expected to be a 32 or 64 character hexadecimal string.
func (c SessionConfig) KeyBytes() (result []byte) {
	result, err := hex.DecodeString(c.Key)
	if err != nil {
		log.Fatalf("Invalid session key: %v", err)
	}
	return
}

// EncKeyBytes returns the session encryption key as a byte array.
// The key is expected to be a 32 or 64 character hexadecimal string.
func (c SessionConfig) EncKeyBytes() (result []byte) {
	result, err := hex.DecodeString(c.EncKey)
	if err != nil {
		log.Fatalf("Invalid session encryption key: %v", err)
	}
	return
}
