package k3sscripts

import (
	"fmt"
	"os"
	"os/exec"
)

func (ctx *Context) StatusK3Pods(s, v, p bool, namespace string) K3sError {
	fmt.Printf("Status for %s:\n", namespace)
	if os.Geteuid() != 0 {
		return NewUnauthorizedError("You must run this script as root")
	}

	if !CheckK3sInstalled() {
		return NewK3sNotInstalledError("k3s is not installed or not running. Please ensure the cluster is deployed.")
	}

	var cmd *exec.Cmd

	if p {
		fmt.Println("\nPODS:")
		fmt.Println("-----")
		cmd = exec.Command("k3s", "kubectl", "get", "pods", "-n", namespace)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return NewK3sNotRunningError(err.Error())
		}
	}

	if s {
		fmt.Println("\nSVCS:")
		fmt.Println("-----")
		cmd = exec.Command("k3s", "kubectl", "get", "services", "-n", namespace)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return NewK3sNotRunningError(err.Error())
		}
	}

	if v {
		fmt.Println("\nPVCS:")
		fmt.Println("-----")
		cmd = exec.Command("k3s", "kubectl", "get", "pvc", "-n", namespace)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return NewK3sNotRunningError(err.Error())
		}
	}

	return nil
}
