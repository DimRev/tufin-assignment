package k3sscripts

import (
	"fmt"
	"os"
	"os/exec"
)

func (ctx *Context) DeployK3sPods() K3sError {
	if !CheckK3sInstalled() {
		return NewK3sNotInstalledError("k3s is not installed or not running. Please ensure the cluster is deployed.")
	}

	if os.Geteuid() != 0 {
		return NewUnauthorizedError("You must run this script as root")
	}

	err := ctx.GenerateManifests()
	defer ctx.CleanupTempFiles()
	if err != nil {
		return err
	}

	for _, manifestPath := range ctx.tempFiles {
		fmt.Printf("Deploying manifest: %s\n", manifestPath)

		cmd := exec.Command("k3s", "kubectl", "apply", "-f", manifestPath)
		cmd.Stdout = nil
		cmd.Stderr = nil

		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to apply manifest %s: %v\n", manifestPath, err)
			return NewFileWriteError(manifestPath, fmt.Sprintf("failed to apply manifest: %v", err))
		}
	}

	fmt.Println("All manifests applied successfully.")
	return nil
}
