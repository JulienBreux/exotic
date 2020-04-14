package config

import (
	"encoding/json"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config represents the configuration structure
type Config struct {
	Debug bool `default:"false"`

	LoggerLevel  string `default:"info" envconfig:"logger_level"`
	LoggerFormat string `default:"json" envconfig:"logger_format"`
}

// JSON exports configuration as JSON
func (c *Config) JSON() []byte {
	b, _ := json.Marshal(c)
	return b
}

// New returns the configuration
func New() (*Config, error) {
	var c Config

	if err := envconfig.Process("", &c); err != nil {
		return nil, errors.Wrap(err, "unable to read configuration")
	}

	return &c, nil
}
