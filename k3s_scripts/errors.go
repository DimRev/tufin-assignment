package k3sscripts

import "fmt"

type K3sError interface {
	Error() string
	K3sError() string
}

type UnauthorizedError struct {
	Message string
}

func (e UnauthorizedError) Error() string {
	return fmt.Sprintf("Unauthorized: %s", e.Message)
}

func (e UnauthorizedError) K3sError() string {
	return fmt.Sprintf("Unauthorized: %s", e.Message)
}

type K3sNotInstalledError struct {
	Message string
}

func (e K3sNotInstalledError) Error() string {
	return fmt.Sprintf("K3s not installed: %s", e.Message)
}

func (e K3sNotInstalledError) K3sError() string {
	return fmt.Sprintf("K3s not installed: %s", e.Message)
}

type K3sNotRunningError struct {
	Message string
}

func (e K3sNotRunningError) Error() string {
	return fmt.Sprintf("K3s not running: %s", e.Message)
}

func (e K3sNotRunningError) K3sError() string {
	return fmt.Sprintf("K3s not running: %s", e.Message)
}

type K3sClusterNotReadyError struct {
	Message string
}

func (e K3sClusterNotReadyError) Error() string {
	return fmt.Sprintf("K3s cluster not ready: %s", e.Message)
}

func (e K3sClusterNotReadyError) K3sError() string {
	return fmt.Sprintf("K3s cluster not ready: %s", e.Message)
}
