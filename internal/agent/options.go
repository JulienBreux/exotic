package agent

import (
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/julienbreux/exotic/internal/logger"
	"github.com/julienbreux/exotic/pkg/path"
)

// ConfigPrefix define configuration prefix
var ConfigPrefix = "agent"

// Option func
type Option func(*Options)

// Options represents the list of options
type Options struct {
	Logger logger.Logger `ignored:"true" json:"-"`

	KubeInCluster bool      `default:"true" envconfig:"in_cluster"`
	KubeFile      path.Path `default:"" envconfig:"file"`

	Frequency time.Duration `default:"3s" envconfig:"frequency"`

	Namespaces []string `default:"default" envconfig:"namespaces"`
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

// KubeInCluster option
func KubeInCluster(v bool) Option {
	return func(o *Options) {
		o.KubeInCluster = v
	}
}

// KubeFile option
func KubeFile(v path.Path) Option {
	return func(o *Options) {
		o.KubeFile = v
	}
}

// Frequency option
func Frequency(v time.Duration) Option {
	return func(o *Options) {
		o.Frequency = v
	}
}

// Namespaces option
func Namespaces(v []string) Option {
	return func(o *Options) {
		o.Namespaces = v
	}
}
