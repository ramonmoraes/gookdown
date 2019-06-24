package main

import (
	"flag"

	gookdown "github.com/ramonmoraes/gookDown/cmd"
)

func main() {
	summaryPath := flag.String("input", "summary.md", "File that contains the file paths to be compiled")
	outputPath := flag.String("output", "README.md", "File path to the output file")
	flag.Parse()
	gookdown.Compile(*summaryPath, *outputPath)
}
