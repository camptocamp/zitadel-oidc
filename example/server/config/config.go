package config

import (
	"os"
	"strings"
)

const (
	// default port for the http server to run
	DefaultIssuerPort = "9998"
)

type Config struct {
	Port         string
	RedirectURIs []string
	UsersFile    string
	Issuer       string
	DevMode      bool
}

// FromEnvVars loads configuration parameters from environment variables.
// If there is no such variable defined, then use default values.
func FromEnvVars(defaults *Config) *Config {
	if defaults == nil {
		defaults = &Config{}
	}
	cfg := &Config{
		Port:         defaults.Port,
		RedirectURIs: defaults.RedirectURIs,
		UsersFile:    defaults.UsersFile,
		Issuer:       defaults.Issuer,
		DevMode:      defaults.DevMode,
	}
	if value, ok := os.LookupEnv("PORT"); ok {
		cfg.Port = value
	}
	if value, ok := os.LookupEnv("USERS_FILE"); ok {
		cfg.UsersFile = value
	}
	if value, ok := os.LookupEnv("REDIRECT_URIS"); ok {
		cfg.RedirectURIs = strings.Split(value, ",")
	}
	if value, ok := os.LookupEnv("ISSUER"); ok {
		cfg.Issuer = value
	}
	if value, ok := os.LookupEnv("DEV_MODE"); ok {
		cfg.DevMode = value == "true"
	}
	return cfg
}
