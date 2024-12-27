package k3sscripts

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func (ctx *Context) DeployK3sCluster() K3sError {
	err := CheckRootUser()
	if err != nil {
		return err
	}

	fmt.Println("Checking if k3s is already installed...")

	if CheckK3sInstalled() {
		fmt.Println("k3s already installed.")
	} else {
		fmt.Println("Starting k3s cluster deployment...")
		cmd := exec.Command("sh", "-c", "curl -sfL https://get.k3s.io | sh -")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return NewK3sNotInstalledError(err.Error())
		}
	}

	fmt.Println("Waiting for the k3s cluster to become ready...")

	if err := ctx.waitForClusterReady(); err != nil {
		return NewK3sNotRunningError(err.Error())
	}

	fmt.Println("k3s cluster deployed and running successfully.")
	return nil
}

func (ctx *Context) waitForClusterReady() error {
	const (
		maxRetries   = 20
		retryDelay   = 5 * time.Second
		successCheck = "Ready"
	)

	for i := 0; i < maxRetries; i++ {
		output, err := exec.Command("k3s", "kubectl", "get", "nodes").Output()
		if err != nil {
			fmt.Printf("Retry %d/%d: Cluster not ready yet...\n", i+1, maxRetries)
		} else if strings.Contains(string(output), successCheck) {
			fmt.Println("Cluster nodes:")
			fmt.Println(strings.TrimSpace(string(output)))
			return nil
		}

		time.Sleep(retryDelay)
	}

	return NewK3sClusterNotReadyError(fmt.Sprintf("Cluster not ready after %d seconds", maxRetries*retryDelay))
}
