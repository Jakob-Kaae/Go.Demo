package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port              string `envconfig:"PORT" default:":4000"`
	DatabaseName      string `envconfig:"DATABASE_NAME" default:"goth.db"`
	SessionCookieName string `envconfig:"SESSION_COOKIE_NAME" default:"session"`
}

func loadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// MustLoadConfig loads and returns the application configuration.
// It calls loadConfig and will panic if loading or validation fails.
// This is a convenience for initialization (for example in main or tests)
// where a missing or invalid configuration should terminate the program.
// Call loadConfig directly if you need to handle errors instead of panicking.
func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
