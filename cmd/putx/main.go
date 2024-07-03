package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CyInnove/putx/config"
	"github.com/CyInnove/putx/internal/runner"
	"github.com/CyInnove/putx/pkg/utils"
)

func main() {
	// Get flags from config package
	options := config.GetFlags()

	// Validate flags
	if options.FilePath == "" {
		fmt.Println("Error: -l flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if !utils.FileExists(options.FilePath) {
		fmt.Printf("Error: File %s does not exist\n", options.FilePath)
		os.Exit(1)
	}

	runner.Run(options)
}
