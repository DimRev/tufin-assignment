package k3sscripts

import (
	"fmt"
	"os/exec"
)

func (ctx *Context) DeployK3sPodsSlim() K3sError {
	if !CheckK3sInstalled() {
		return NewK3sNotInstalledError("k3s is not installed or not running. Please ensure the cluster is deployed.")
	}

	err := CheckRootUser()
	if err != nil {
		return err
	}

	err = ctx.GenerateManifests()
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

func (ctx *Context) DeployK3sPodsHelm() K3sError {
	if !CheckK3sInstalled() {
		return NewK3sNotInstalledError("k3s is not installed or not running. Please ensure the cluster is deployed.")
	}

	err := CheckRootUser()
	if err != nil {
		return err
	}

	if !CheckHelmInstalled() {
		return NewHelmInstallError("helm is not installed. Please ensure helm is installed.")
	}

	err = ctx.GenerateHelmChart()
	defer ctx.CleanupTempFiles()
	if err != nil {
		return err
	}

	if len(ctx.tempFiles) != 1 {
		return NewYAMLFilesNotFound(
			"wordpress-sql-1.0.0.tgz",
			fmt.Sprintf("expected 1 file, found %d", len(ctx.tempFiles)),
		)
	}

	helmChartPath := ctx.tempFiles[0]

	fmt.Printf("Deploying helm chart: %s\n", helmChartPath)

	return nil
}
