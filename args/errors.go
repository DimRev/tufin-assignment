package args

import (
	"fmt"
	"strings"
)

type ArgErrors interface {
	Error() string
	ArgErrors() string
	Code() int
}

type GenericArgErrors struct {
	Message    string
	StatusCode int
}

func (e GenericArgErrors) Error() string {
	return e.Message
}

func (e GenericArgErrors) ArgErrors() string {
	return e.Message
}

func (e GenericArgErrors) Code() int {
	return e.StatusCode
}

type CommandNotFoundError struct {
	GenericArgErrors
}

type InvalidFlagError struct {
	GenericArgErrors
}

type UnknownArgsError struct {
	GenericArgErrors
}

func NewCommandNotFoundError(commandName string) ArgErrors {
	return &CommandNotFoundError{
		GenericArgErrors{
			Message:    fmt.Sprintf("ArgsCommandNotFound: %s", commandName),
			StatusCode: 101,
		},
	}
}

func NewInvalidFlagError(flagName string, command CommandName) ArgErrors {
	return &InvalidFlagError{
		GenericArgErrors{
			Message:    fmt.Sprintf("ArgsInvalidFlag: %s, for command '%s'", flagName, command),
			StatusCode: 102,
		},
	}
}

func NewUnknownArgsError(args []string, command CommandName) ArgErrors {
	return &UnknownArgsError{
		GenericArgErrors{
			Message:    fmt.Sprintf("ArgsUnknownArgs: %s, for command '%s'", strings.Join(args, ", "), command),
			StatusCode: 103,
		},
	}
}
