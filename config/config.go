package config

import (
	"flag"

	"github.com/CyInnove/putx/internal/runner"
)

// GetFlags defines and parses the command-line flags.
func GetFlags() *runner.Options {
	options := &runner.Options{}
	flag.StringVar(&options.FilePath, "l", "", "File containing base URLs")
	flag.StringVar(&options.OutputFile, "o", "", "Output file")
	flag.Parse()
	return options
}
