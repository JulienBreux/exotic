package logger

import "github.com/kelseyhightower/envconfig"

// ConfigPrefix define configuration prefix
var ConfigPrefix = "logger"

// Option func
type Option func(*Options)

// Options represents the list of options
type Options struct {
	Level  string `default:"info" envconfig:"level"`
	Format string `default:"json" envconfig:"format"`

	InstName    string `default:"unknown" envconfig:"inst_name"`
	InstVersion string `default:"unknown" envconfig:"inst_version"`

	Debug bool `default:"false" envconfig:"debug"`
}

func newOptions(opts ...Option) (*Options, error) {
	opt := Options{}
	err := envconfig.Process(ConfigPrefix, &opt)
	if err != nil {
		return nil, err
	}

	for _, o := range opts {
		o(&opt)
	}

	if opt.Level == "" {
		opt.Level = "info"
	}
	if opt.Format == "" {
		opt.Format = "text"
	}

	return &opt, nil
}

// Level option
func Level(v string) Option {
	return func(o *Options) {
		o.Level = v
	}
}

// Format option
func Format(v string) Option {
	return func(o *Options) {
		o.Format = v
	}
}

// InstName option
func InstName(v string) Option {
	return func(o *Options) {
		o.InstName = v
	}
}

// InstVersion option
func InstVersion(v string) Option {
	return func(o *Options) {
		o.InstVersion = v
	}
}

// Debug option
func Debug(v bool) Option {
	return func(o *Options) {
		o.Debug = v
	}
}
