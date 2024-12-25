package k3sscripts

import (
	"fmt"
	"os"
	"os/exec"
)

func (ctx *Context) StatusK3Pods() K3sError {
	if os.Geteuid() != 0 {
		return NewUnauthorizedError("You must run this script as root")
	}

	if !CheckK3sInstalled() {
		return NewK3sNotInstalledError("k3s is not installed or not running. Please ensure the cluster is deployed.")
	}

	fmt.Println("PODS:")
	fmt.Println("-----")
	cmd := exec.Command("k3s", "kubectl", "get", "pods", "-n", "default")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return NewK3sNotRunningError(err.Error())
	}

	fmt.Println("\nSVCS:")
	fmt.Println("-----")
	cmd = exec.Command("k3s", "kubectl", "get", "services", "-n", "default")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return NewK3sNotRunningError(err.Error())
	}

	fmt.Println("\nPVCS:")
	fmt.Println("-----")
	cmd = exec.Command("k3s", "kubectl", "get", "pvc", "-n", "default")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return NewK3sNotRunningError(err.Error())
	}

	return nil
}
