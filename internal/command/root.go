package command

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/julienbreux/exotic/internal/command/serve"
	"github.com/julienbreux/exotic/internal/command/version"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exotic",
	Short: "Secret project",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(buildVersion, buildCommit, buildDate string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	rootCmd.Version = buildVersion

	// Commands
	rootCmd.AddCommand(version.New(buildVersion, buildCommit, buildDate))
	rootCmd.AddCommand(serve.New())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
