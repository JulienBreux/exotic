package manager

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"github.com/julienbreux/exotic/internal/config"
)

// Manager represents the manager structure
type Manager struct {
	ctx  context.Context
	conf *config.Config
	log  zerolog.Logger

	cmpts []Component
}

// New returns a manager instance
func New(ctx context.Context, conf *config.Config, log zerolog.Logger) *Manager {
	return &Manager{
		ctx:  ctx,
		conf: conf,
		log:  log.With().Str("process", "manager").Logger(),

		cmpts: []Component{},
	}
}

// Add adds a new components in the manager
func (m *Manager) Add(c Component) error {
	m.cmpts = append(m.cmpts, c)

	return nil
}

// Start starts the manager
func (m *Manager) Start() error {

	// Manage service interuption
	ctx, cancel := context.WithCancel(m.ctx)
	go func() {
		// Stop :(
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		m.log.Info().Str("reason", (<-ch).String()).Msg("Stopping...")
		cancel()

		if err := m.Stop(); err != nil {
			m.log.Error().Err(err).Msg("Stopping error")
		}
		m.log.Info().Msg("Stopped")
	}()

	// Manage components life cycle
	m.log.Debug().Msg("Starting...")
	eg, ctx := errgroup.WithContext(ctx)
	for _, c := range m.cmpts {
		eg.Go(func() error {
			return c.Start(ctx)
		})
	}

	m.log.Info().Msg("Started")

	if err := eg.Wait(); err != nil {
		m.log.Error().Err(err).Msg("Starting error")
	}

	return nil
}

// Stop stops the  manager
func (m *Manager) Stop() error {
	for _, c := range m.cmpts {
		_ = c.Stop()
	}

	return nil
}
