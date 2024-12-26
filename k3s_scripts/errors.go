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

func NewUnauthorizedError(message string) K3sError {
	return &UnauthorizedError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3UnauthorizedError: %s", message),
			StatusCode: 201,
		},
	}
}

func NewK3sNotInstalledError(message string) K3sError {
	return &K3sNotInstalledError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3InstallError: %s", message),
			StatusCode: 202,
		},
	}
}

func NewK3sNotRunningError(message string) K3sError {
	return &K3sNotRunningError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3NotRunning: %s", message),
			StatusCode: 203,
		},
	}
}

func NewK3sClusterNotReadyError(message string) K3sError {
	return &K3sClusterNotReadyError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3ClusterNotReady: %s", message),
			StatusCode: 204,
		},
	}
}

func NewYAMLFilesNotFound(fileName, message string) K3sError {
	return &YAMLFilesNotFound{
		GenericK3sError{
			Message:    fmt.Sprintf("K3FileNotFound: %s", fileName),
			StatusCode: 205,
		},
		fileName,
	}
}

func NewYAMLUnmarshalError(message string) K3sError {
	return &YAMLUnmarshalError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3UnmarshalError: %s", message),
			StatusCode: 206,
		},
	}
}

func NewDirCreationError(fileName, message string) K3sError {
	return &DirCreationError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3CreateDirError: %s, Details: %s", fileName, message),
			StatusCode: 207,
		},
		fileName,
	}
}

func NewFileReadError(fileName, message string) K3sError {
	return &FileReadError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3ReadFileError: %s, Details: %s", fileName, message),
			StatusCode: 208,
		},
		fileName,
	}
}

func NewFileRenderError(fileName, message string) K3sError {
	return &FileRenderError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3RenderFileError: %s, Details: %s", fileName, message),
			StatusCode: 209,
		},
		fileName,
	}
}

func NewFileWriteError(fileName, message string) K3sError {
	return &FileWriteError{
		GenericK3sError{Message: fmt.Sprintf("K3WriteFileError: %s, Details: %s", fileName, message),
			StatusCode: 210,
		},
		fileName,
	}
}

func NewYAMLMarshalError(message string) K3sError {
	return &YAMLMarshalError{
		GenericK3sError{
			Message:    fmt.Sprintf("K3MarshalError: %s", message),
			StatusCode: 211,
		},
	}
}

func NewValuesReplacementError(message string) K3sError {
	return &ValuesReplacementError{
		GenericK3sError{
			Message: fmt.Sprintf("K3ValuesReplacementError: %s", message),
		},
	}
}
