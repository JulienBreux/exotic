package master

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/julienbreux/exotic/internal/logger"
	"github.com/julienbreux/exotic/internal/manager"
	"github.com/julienbreux/exotic/internal/master"
)

// New creates a new command instance
func New(lgr logger.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "master",
		Short: "Start the Exotic master",
		Run:   run(lgr),
	}
}

func run(lgr logger.Logger) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Master
		mstr, _ := master.New(master.Logger(lgr))

		// Manager
		m, _ := manager.New(
			manager.Logger(lgr),
			manager.Context(context.Background()),
			manager.AddComponent(mstr),
		)
		if err := m.Run(); err != nil {
			os.Exit(1)
		}
	}
}
