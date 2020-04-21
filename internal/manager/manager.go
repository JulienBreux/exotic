package manager

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

const defaultProcess = "manager"

type manager struct {
	opt *Options
	lgr zerolog.Logger
	ctx context.Context

	cmpts []Component
}

// New creates a new manager instance
func New(opts ...Option) (Manager, error) {
	opt, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	return &manager{
		opt: opt,
		lgr: opt.Logger.Log().With().Str("process", defaultProcess).Logger(),
		ctx: opt.Context,

		cmpts: opt.Components,
	}, nil
}

// Options returns the list of options
func (m *manager) Options() *Options {
	return m.opt
}

// Run ran the manager
func (m *manager) Run() error {
	// Manage service interuption
	ctx, cancel := context.WithCancel(m.ctx)
	go m.interrupt(ctx, cancel)

	// Manage components life cycle
	m.lgr.Debug().Msg("Starting components...")
	eg, ctx := errgroup.WithContext(ctx)

	for _, c := range m.cmpts {
		eg.Go(func() error {
			return c.Start(ctx)
		})
	}

	m.lgr.Info().Msg("All component started")

	if err := eg.Wait(); err != nil {
		m.lgr.Error().Err(err).Msg("Starting error")
		return errors.Wrap(err, "Unable to start components")
	}

	return nil
}

// Stop stops the  manager
func (m *manager) Stop() (err error) {
	for _, c := range m.cmpts {
		if err = c.Stop(); err != nil {
			err = errors.Wrap(err, "unable to stop component")
			m.lgr.Warn().Err(err).Msgf("Unable to stop component %s", "")
		}
	}

	if err == nil {
		m.lgr.Info().Msg("All components stopped")
	}

	return
}

func (m *manager) interrupt(ctx context.Context, c context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	m.lgr.Info().Str("reason", (<-ch).String()).Msg("Stopping...")
	c()
	if err := m.Stop(); err != nil {
		m.lgr.Error().Err(err).Msg("Stopping error")
	}
	m.lgr.Info().Msg("Stopped")
}
