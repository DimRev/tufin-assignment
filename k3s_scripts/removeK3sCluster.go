package k3sscripts

import (
	"fmt"
	"os"
	"os/exec"
)

func RemoveK3sCluster() K3sError {
	fmt.Println("Removing k3s cluster...")

	if !CheckK3sInstalled() {
		return K3sNotInstalledError{Message: "k3s not installed"}
	}

	if os.Geteuid() != 0 {
		return UnauthorizedError{Message: "You must run this script as root"}
	}

	uninstallScript := "/usr/local/bin/k3s-uninstall.sh"
	if _, err := os.Stat(uninstallScript); os.IsNotExist(err) {
		return K3sNotInstalledError{Message: "Uninstall script not found. Is k3s installed?"}
	}

	cmd := exec.Command("sh", uninstallScript)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running k3s uninstall script...")
	if err := cmd.Run(); err != nil {
		return K3sNotRunningError{Message: fmt.Sprintf("Failed to remove k3s: %v", err)}
	}

	fmt.Println("k3s cluster removed successfully.")
	return nil
}
