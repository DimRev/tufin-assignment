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

func main() {
	var k3ssCtx = k3sscripts.NewContext(manifests)
	var commandMap = newCommandMap(k3ssCtx)

	inputArgs := os.Args[1:]
	err := args.ParseArgs(inputArgs, commandMap)
	if err != nil {
		switch err := err.(type) {
		case args.ArgErrors:
			fmt.Printf("ArgsError: %v (Code: %d)\n", err, err.Code())
			os.Exit(err.Code())
		case k3sscripts.K3sError:
			fmt.Printf("K3sError: %v (Code: %d)\n", err, err.Code())
			os.Exit(err.Code())
		default:
			fmt.Printf("UnknownError: %v (Code: %d)\n", err, 1)
			os.Exit(1)
		}
	}
}

func newCommandMap(k3ssCtx *k3sscripts.Context) map[args.CommandName]args.ExecutionFunc {
	return map[args.CommandName]args.ExecutionFunc{
		args.GlobalCommand: func(flags map[string]string) error {
			args.HelpPrint(args.GlobalCommand)
			return nil
		},
		args.ClusterCommand: func(flags map[string]string) error {
			err := k3ssCtx.DeployK3sCluster()
			if err != nil {
				return err
			}
			return nil
		},
		args.DeployCommand: func(flags map[string]string) error {
			err := k3ssCtx.DeployK3sPods()
			if err != nil {
				return err
			}

			return nil
		},
		args.StatusCommand: func(flags map[string]string) error {
			showService := flags["--service"] == "true"
			showVolume := flags["--volume"] == "true"
			showPod := flags["--pod"] == "true"
			showNamespace := flags["--namespace"]

			if showNamespace == "" {
				showNamespace = "default"
			}

			if !showService && !showVolume && !showPod {
				showService = true
				showVolume = true
				showPod = true
			}

			err := k3ssCtx.StatusK3Pods(showService, showVolume, showPod, showNamespace)
			if err != nil {
				return err
			}
			return nil
		},
		args.RemoveCommand: func(flags map[string]string) error {
			err := k3ssCtx.RemoveK3sCluster()
			if err != nil {
				return err
			}
			return nil
		},
	}
}
