package version

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	app     = "Exotic"
	author  = "Julien BREUX <julien.breux@gmail.com>"
	website = "https://github.com/JulienBreux/exotic/"
)

// New creates a new version instance
func New(version, commit, date string) *cobra.Command {
	info := make(map[string]string, 5)
	info["version"] = version
	info["commit"] = commit
	info["date"] = date
	info["author"] = author
	info["website"] = website

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version/build info",
		Long:  "Show the version and the build information",
		Run:   versionRun(info),
	}

	cmd.PersistentFlags().Bool("only", false, "print only version")

	return cmd
}

func versionRun(info map[string]string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.PersistentFlags().GetBool("only"); v {
			fmt.Println(info["version"])
			return
		}
		printTitle()
		for k, v := range info {
			printKeyValue(k, v)
		}
		fmt.Println("")
	}
}

func printKeyValue(key, value string) {
	fmt.Println(
		aurora.Cyan(
			fmt.Sprintf("%7v", strings.ToTitle(key)),
		),
		aurora.White(value),
	)
}

func printTitle() {
	s := aurora.Green(strings.Repeat("~", 17))
	fmt.Printf(
		"\n%s 🌴 %s™ %s\n\n",
		s,
		aurora.White(app),
		s,
	)
}
