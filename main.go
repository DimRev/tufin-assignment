package main

import (
	"fmt"
	"os"

	"github.com/DimRev/tufin-assignment/args"
	k3sscripts "github.com/DimRev/tufin-assignment/k3s_scripts"
)

var commandMap = map[args.CommandName]func() error{
	args.GlobalCommand: func() error {
		args.HelpPrint(args.GlobalCommand)
		return nil
	},
	args.ClusterCommand: func() error {
		err := k3sscripts.DeployK3sCluster()
		if err != nil {
			return err
		}
		return nil
	},
	args.DeployCommand: func() error {
		fmt.Println("Deploying two pods: MySQL and WordPress")
		return nil
	},
	args.StatusCommand: func() error {
		fmt.Println("Printing the status table of pods in the default namespace")
		return nil
	},
	args.RemoveCommand: func() error {
		err := k3sscripts.RemoveK3sCluster()
		if err != nil {
			return err
		}
		return nil
	},
}

func main() {
	// Pass command-line arguments (excluding program name) to ParseArgs
	inputArgs := os.Args[1:]
	err := args.ParseArgs(inputArgs, commandMap)
	if err != nil {
		fmt.Println(err)
	}
}
