package master

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/julienbreux/exotic/internal/logger"
)

// ConfigPrefix define configuration prefix
var ConfigPrefix = "master"

// Option func
type Option func(*Options)

// Options represents the list of options
type Options struct {
	Logger logger.Logger `ignored:"true" json:"-"`

	Port int `default:"8181" envconfig:"port"`
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

	// Set default values
	if opt.Logger == nil {
		if opt.Logger == nil {
			opt.Logger, _ = logger.New()
		}
	}

	return &opt, nil
}

// Logger option
func Logger(v logger.Logger) Option {
	return func(o *Options) {
		o.Logger = v
	}
}

// Port option
func Port(v int) Option {
	return func(o *Options) {
		o.Port = v
	}
}
