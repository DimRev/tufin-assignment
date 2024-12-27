package k3sscripts

import (
	"fmt"
	"os"
	"os/exec"
)

func (ctx *Context) RemoveK3sCluster() K3sError {
	fmt.Println("Removing k3s cluster...")

	if !CheckK3sInstalled() {
		return NewK3sNotInstalledError("K3 is not installed")
	}

	err := CheckRootUser()
	if err != nil {
		return err
	}

	uninstallScript := "/usr/local/bin/k3s-uninstall.sh"
	if _, err := os.Stat(uninstallScript); os.IsNotExist(err) {
		return NewK3sNotInstalledError("Uninstall script not found. Is k3s installed?")
	}

	cmd := exec.Command("sh", uninstallScript)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running k3s uninstall script...")
	if err := cmd.Run(); err != nil {
		return NewK3sNotRunningError(fmt.Sprintf("Failed to remove k3s: %v", err))
	}

	fmt.Println("k3s cluster removed successfully.")
	return nil
}
