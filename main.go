package main

import (
	"fmt"
	"os"

	"github.com/DimRev/tufin-assignment/args"
)

var commandMap = map[args.CommandName]func(){
	args.GlobalCommand: func() {
		args.HelpPrint(args.GlobalCommand)
	},
	args.ClusterCommand: func() {
		fmt.Println("Deploying a k3s cluster")
	},
	args.DeployCommand: func() {
		fmt.Println("Deploying two pods: MySQL and WordPress")
	},
	args.StatusCommand: func() {
		fmt.Println("Printing the status table of pods in the default namespace")
	},
}

func main() {
	// Pass command-line arguments (excluding program name) to ParseArgs
	inputArgs := os.Args[1:]
	args.ParseArgs(inputArgs, commandMap)
}
