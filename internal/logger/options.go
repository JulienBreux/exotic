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

// DefaultFields fields option
// func DefaultFields(fields ...Field) Option {
// 	return func(o *Options) {
// 		o.DefaultFields = fields
// 	}
// }

// Level option
func Level(level string) Option {
	return func(o *Options) {
		o.Level = level
	}
}

// Format option
func Format(format string) Option {
	return func(o *Options) {
		o.Format = format
	}
}

// InstName option
func InstName(instName string) Option {
	return func(o *Options) {
		o.InstName = instName
	}
}

// InstVersion option
func InstVersion(instVersion string) Option {
	return func(o *Options) {
		o.InstVersion = instVersion
	}
}

// Debug option
func Debug(debug bool) Option {
	return func(o *Options) {
		o.Debug = debug
	}
}
