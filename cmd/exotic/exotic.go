package main

import (
	"github.com/julienbreux/exotic/internal/command"
)

var (
	buildVersion = "dev"
	buildCommit  = "dev"
	buildDate    = "n/a"
)

func main() {
	command.Execute(buildVersion, buildCommit, buildDate)
}
