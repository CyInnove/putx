package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CyInnove/putx/internal/runner"
	"github.com/CyInnove/putx/pkg/utils"
)

func main() {
	// Define command-line flags
	outputFile := flag.String("o", "", "Output file")
	inputFile := flag.String("l", "", "File containing base URLs")
	flag.Parse()

	// Validate flags
	if *inputFile == "" {
		fmt.Println("Error: -l flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if !utils.FileExists(*inputFile) {
		fmt.Printf("Error: File %s does not exist\n", *inputFile)
		os.Exit(1)
	}

	runner.Run(*inputFile, *outputFile)
}
