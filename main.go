package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/DimRev/tufin-assignment/args"
	k3sscripts "github.com/DimRev/tufin-assignment/k3s_scripts"
)

//go:embed manifests/*
var manifests embed.FS
var k3ssCtx = k3sscripts.NewContext(manifests)

var commandMap = map[args.CommandName]func() error{
	args.GlobalCommand: func() error {
		args.HelpPrint(args.GlobalCommand)
		return nil
	},
	args.ClusterCommand: func() error {
		err := k3ssCtx.DeployK3sCluster()
		if err != nil {
			return err
		}
		return nil
	},
	args.DeployCommand: func() error {
		err := k3ssCtx.DeployK3sPods()
		if err != nil {
			return err
		}

		return nil
	},
	args.StatusCommand: func() error {
		fmt.Println("Printing the status table of pods in the default namespace")
		return nil
	},
	args.RemoveCommand: func() error {
		err := k3ssCtx.RemoveK3sCluster()
		if err != nil {
			return err
		}
		return nil
	},
}

func main() {
	inputArgs := os.Args[1:]
	err := args.ParseArgs(inputArgs, commandMap)
	if err != nil {
		fmt.Println(err)
	}
}
