package k3sscripts

import "os/exec"

func CheckK3sInstalled() bool {
	_, err := exec.LookPath("k3s")
	return err == nil
}
