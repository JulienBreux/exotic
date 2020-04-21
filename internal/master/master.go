package master

import (
	"context"

	"github.com/rs/zerolog"
)

const defaultProcess = "master"

type master struct {
	opt *Options

	lgr zerolog.Logger
}

// New creates a new master instance
func New(opts ...Option) (Master, error) {
	opt, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	return &master{
		opt: opt,
		lgr: opt.Logger.Log().With().Str("process", defaultProcess).Logger(),
	}, nil
}

// Options returns the list of options
func (m *master) Options() *Options {
	return m.opt
}

// Start starts the agent
func (m *master) Start(ctx context.Context) (err error) {
	m.lgr.Info().Msg("Started")

	// FIXME: Write master code part
	<-ctx.Done()

	return
}

// Stop stops the agent
func (m *master) Stop() (err error) {
	m.lgr.Info().Msg("Stopped")
	return
}
