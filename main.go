package main

import (
	"os"

	"github.com/DimRev/tufin-assignment/args"
)

func main() {
	// Pass command-line arguments (excluding program name) to ParseArgs
	inputArgs := os.Args[1:]
	args.ParseArgs(inputArgs)
}
