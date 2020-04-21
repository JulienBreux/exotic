package agent

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/julienbreux/exotic/internal/agent"
	"github.com/julienbreux/exotic/internal/logger"
	"github.com/julienbreux/exotic/internal/manager"
)

// New creates a new command instance
func New(lgr logger.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "agent",
		Short: "Start the Exotic agent",
		Run:   run(lgr),
	}
}

func run(lgr logger.Logger) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Agent
		agt, err := agent.New(agent.Logger(lgr))
		if err != nil {
			l := lgr.Log()
			l.Fatal().Msg(err.Error())
			os.Exit(1)
		}

		// Manager
		m, _ := manager.New(
			manager.Logger(lgr),
			manager.Context(context.Background()),
			manager.AddComponent(agt),
		)
		if err := m.Run(); err != nil {
			os.Exit(1)
		}
	}
}
