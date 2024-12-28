package k3sscripts

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func (ctx *Context) RemoveHelmChart() K3sError {
	if !CheckHelmInstalled() {
		return nil
	}

	// Command to list installed Helm charts
	cmd := exec.Command("helm", "list", "--kubeconfig=/etc/rancher/k3s/k3s.yaml", "--short")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		return NewHelmUninstallError(fmt.Sprintf("failed to list installed Helm charts: %v", err))
	}

	// Check if the specific chart is in the list of installed charts
	installedCharts := strings.Split(stdout.String(), "\n")
	for _, chart := range installedCharts {
		chart = strings.TrimSpace(chart)
		if chart == "my-wordpress-sql" {
			// Chart exists, uninstall it
			fmt.Println("Helm chart 'my-wordpress-sql' is installed. Removing it...")

			uninstallCmd := exec.Command("helm", "uninstall", "my-wordpress-sql", "--kubeconfig=/etc/rancher/k3s/k3s.yaml")
			uninstallCmd.Stdout = nil
			uninstallCmd.Stderr = nil

			if err := uninstallCmd.Run(); err != nil {
				return NewHelmUninstallError(fmt.Sprintf("failed to uninstall Helm chart: %v", err))
			}

			fmt.Println("Helm chart 'my-wordpress-sql' has been removed successfully.")
			return nil
		}
	}

	// Chart does not exist
	fmt.Println("Helm chart 'my-wordpress-sql' is not installed. No action required.")
	return nil
}
