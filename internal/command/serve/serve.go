package serve

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/julienbreux/exotic/internal/config"
	"github.com/julienbreux/exotic/internal/manager"
)

// New creates a new command instance
func New() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the Exotic server",
		Run:   serveRun,
	}
}

func serveRun(cmd *cobra.Command, args []string) {
	lgr := log.With().Logger()
	// Init configuration
	c, err := config.New()
	if err != nil {
		lgr.Error().Str("process", "global").Err(err).Msg("Unable to load environment variable")
		os.Exit(1)
	}
	if c.LoggerFormat == "text" {
		lgr = lgr.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}

	// Init logger
	level, _ := zerolog.ParseLevel(c.LoggerLevel)
	zerolog.SetGlobalLevel(level)
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		lgr.Debug().Str("process", "global").Msg("Debug mode enabled")
		lgr.Debug().RawJSON("config", c.JSON()).Str("process", "global").Msg("Environment variables loaded")
	}

	// Init manager
	ctx := context.Background()
	m := manager.New(ctx, c, lgr)
	// TODO: m.Add(exampleComponent.New(c, lgr))
	if err := m.Start(); err != nil {
		os.Exit(1)
	}
}
