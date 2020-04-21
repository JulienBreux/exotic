package client

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/julienbreux/exotic/internal/logger"
	"github.com/julienbreux/exotic/pkg/path"
)

// ConfigPrefix define configuration prefix
var ConfigPrefix = "client"

// Option func
type Option func(*Options)

// Options represents the list of options
type Options struct {
	Logger logger.Logger `ignored:"true" json:"-"`

	InCluster  bool      `default:"true" envconfig:"in_cluster"`
	ConfigPath path.Path `envconfig:"config_path"`

	TimeoutSeconds int64 `default:"2" envconfig:"timeout_seconds"`
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

// TimeoutSeconds option
func TimeoutSeconds(v int64) Option {
	return func(o *Options) {
		o.TimeoutSeconds = v
	}
}
