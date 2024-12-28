package k3sscripts

import (
	"os"
	"os/exec"
)

func CheckK3sInstalled() bool {
	_, err := exec.LookPath("k3s")
	return err == nil
}

func CheckHelmInstalled() bool {
	_, err := exec.LookPath("helm")
	return err == nil
}

func CheckRootUser() K3sError {
	if os.Geteuid() != 0 {
		return NewUnauthorizedError("You must run this script as root")
	}
	return nil
}
