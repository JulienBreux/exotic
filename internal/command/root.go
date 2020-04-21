package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/julienbreux/exotic/internal/command/agent"
	"github.com/julienbreux/exotic/internal/command/master"
	"github.com/julienbreux/exotic/internal/command/version"
	"github.com/julienbreux/exotic/internal/logger"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exotic",
	Short: "Secret project",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(buildVersion, buildCommit, buildDate string) {
	rootCmd.Version = buildVersion

	// Logger
	lgr, err := logger.New(
		logger.InstName(rootCmd.Use), // TODO: check value
		logger.InstVersion(rootCmd.Version),
	)
	if err != nil {
		os.Exit(1)
	}

	// Commands
	rootCmd.AddCommand(version.New(buildVersion, buildCommit, buildDate))
	rootCmd.AddCommand(master.New(lgr))
	rootCmd.AddCommand(agent.New(lgr))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
