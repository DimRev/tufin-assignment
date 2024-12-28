package k3sscripts

import (
	"fmt"
)

type K3sError interface {
	Error() string
	K3sError() string
	Code() int
}

type GenericK3sError struct {
	Message    string
	StatusCode int
}

const (
	UnauthorizedErrorCode       = 201
	K3sNotInstalledErrorCode    = 202
	K3sNotRunningErrorCode      = 203
	K3sClusterNotReadyErrorCode = 204
	YAMLFilesNotFoundErrorCode  = 205
	YAMLUnmarshalErrorCode      = 206
	DirCreationErrorCode        = 207
	FileReadErrorCode           = 208
	FileRenderErrorCode         = 209
	FileWriteErrorCode          = 210
	YAMLMarshalErrorCode        = 211
	ValuesReplacementErrorCode  = 212
	HelmInstallErrorCode        = 213
	HelmUninstallErrorCode      = 214
	HelmDeployErrorCode         = 215
)

func (e GenericK3sError) Error() string {
	return e.Message
}

func (e GenericK3sError) K3sError() string {
	return e.Message
}

func (e GenericK3sError) Code() int {
	return e.StatusCode
}

type UnauthorizedError struct {
	GenericK3sError
}

type K3sNotInstalledError struct {
	GenericK3sError
}

type K3sNotRunningError struct {
	GenericK3sError
}

type K3sClusterNotReadyError struct {
	GenericK3sError
}

type YAMLFilesNotFound struct {
	GenericK3sError
	FileName string
}

type YAMLUnmarshalError struct {
	GenericK3sError
}

type DirCreationError struct {
	GenericK3sError
	DirName string
}

type FileReadError struct {
	GenericK3sError
	FileName string
}

type FileRenderError struct {
	GenericK3sError
	FileName string
}

type FileWriteError struct {
	GenericK3sError
	FileName string
}

type YAMLMarshalError struct {
	GenericK3sError
}

type ValuesReplacementError struct {
	GenericK3sError
}

type HelmInstallError struct {
	GenericK3sError
}

type HelmUninstallError struct {
	GenericK3sError
}

type HelmDeployError struct {
	GenericK3sError
}

func NewUnauthorizedError(message string) K3sError {
	return &UnauthorizedError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3UnauthorizedError: %s", message),
			StatusCode: UnauthorizedErrorCode,
		},
	}
}

func NewK3sNotInstalledError(message string) K3sError {
	return &K3sNotInstalledError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3InstallError: %s", message),
			StatusCode: K3sNotInstalledErrorCode,
		},
	}
}

func NewK3sNotRunningError(message string) K3sError {
	return &K3sNotRunningError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3NotRunning: %s", message),
			StatusCode: K3sNotRunningErrorCode,
		},
	}
}

func NewK3sClusterNotReadyError(message string) K3sError {
	return &K3sClusterNotReadyError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3ClusterNotReady: %s", message),
			StatusCode: K3sClusterNotReadyErrorCode,
		},
	}
}

func NewYAMLFilesNotFound(fileName, message string) K3sError {
	return &YAMLFilesNotFound{
		GenericK3sError{
			Message:    fmt.Sprintf("K3FileNotFound: %s", fileName),
			StatusCode: YAMLFilesNotFoundErrorCode,
		},
		fileName,
	}
}

func NewYAMLUnmarshalError(message string) K3sError {
	return &YAMLUnmarshalError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3UnmarshalError: %s", message),
			StatusCode: YAMLUnmarshalErrorCode,
		},
	}
}

func NewDirCreationError(fileName, message string) K3sError {
	return &DirCreationError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3CreateDirError: %s, Details: %s", fileName, message),
			StatusCode: DirCreationErrorCode,
		},
		fileName,
	}
}

func NewFileReadError(fileName, message string) K3sError {
	return &FileReadError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3ReadFileError: %s, Details: %s", fileName, message),
			StatusCode: FileReadErrorCode,
		},
		fileName,
	}
}

func NewFileRenderError(fileName, message string) K3sError {
	return &FileRenderError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3RenderFileError: %s, Details: %s", fileName, message),
			StatusCode: FileRenderErrorCode,
		},
		fileName,
	}
}

func NewFileWriteError(fileName, message string) K3sError {
	return &FileWriteError{
		GenericK3sError{Message: fmt.Sprintf("K3WriteFileError: %s, Details: %s", fileName, message),
			StatusCode: FileWriteErrorCode,
		},
		fileName,
	}
}

func NewYAMLMarshalError(message string) K3sError {
	return &YAMLMarshalError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3MarshalError: %s", message),
			StatusCode: YAMLMarshalErrorCode,
		},
	}
}

func NewValuesReplacementError(message string) K3sError {
	return &ValuesReplacementError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3ValuesReplacementError: %s", message),
			StatusCode: ValuesReplacementErrorCode,
		},
	}
}

func NewHelmInstallError(message string) K3sError {
	return &HelmInstallError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3HelmInstallError: %s", message),
			StatusCode: HelmInstallErrorCode,
		},
	}
}

func NewHelmUninstallError(message string) K3sError {
	return &HelmUninstallError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3HelmUninstallError: %s", message),
			StatusCode: HelmUninstallErrorCode,
		},
	}
}

func NewHelmDeployError(message string) K3sError {
	return &HelmDeployError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3HelmDeployError: %s", message),
			StatusCode: HelmDeployErrorCode,
		},
	}
}
