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

type YAMLFilesNotFound struct {
	FileName string
	Message  string
}

func (e YAMLFilesNotFound) Error() string {
	return fmt.Sprintf("YAML files %s not found: %s", e.FileName, e.Message)
}

func (e YAMLFilesNotFound) K3sError() string {
	return fmt.Sprintf("YAML files %s not found: %s", e.FileName, e.Message)
}

type YAMLUnmarshalError struct {
	Message string
}

func (e YAMLUnmarshalError) Error() string {
	return fmt.Sprintf("Error unmarshalling YAML: %s", e.Message)
}

func (e YAMLUnmarshalError) K3sError() string {
	return fmt.Sprintf("Error unmarshalling YAML: %s", e.Message)
}

type YAMLMarshalError struct {
	Message string
}

func (e YAMLMarshalError) Error() string {
	return fmt.Sprintf("Error marshalling YAML: %s", e.Message)
}

func (e YAMLMarshalError) K3sError() string {
	return fmt.Sprintf("Error marshalling YAML: %s", e.Message)
}

type ValuesReplacementError struct {
	Message string
}

func (e ValuesReplacementError) Error() string {
	return fmt.Sprintf("Error replacing values: %s", e.Message)
}

func (e ValuesReplacementError) K3sError() string {
	return fmt.Sprintf("Error replacing values: %s", e.Message)
}
