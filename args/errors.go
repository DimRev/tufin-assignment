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

const (
	CommandNotFoundErrorCode = 101
	InvalidFlagErrorCode     = 102
	UnknownArgsErrorCode     = 103
	FlagArgMissingErrorCode  = 104
	FlagCombinationErrorCode = 105
)

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

type FlagArgMissing struct {
	GenericArgErrors
}

type FlagCombinationError struct {
	GenericArgErrors
}

func NewCommandNotFoundError(commandName string) ArgErrors {
	return &CommandNotFoundError{
		GenericArgErrors{
			Message:    fmt.Sprintf("ArgsCommandNotFound: %s", commandName),
			StatusCode: CommandNotFoundErrorCode,
		},
	}
}

func NewInvalidFlagError(flagName string, command CommandName) ArgErrors {
	return &InvalidFlagError{
		GenericArgErrors{
			Message:    fmt.Sprintf("ArgsInvalidFlag: %s, for command '%s'", flagName, command),
			StatusCode: InvalidFlagErrorCode,
		},
	}
}

func NewUnknownArgsError(args []string, command CommandName) ArgErrors {
	return &UnknownArgsError{
		GenericArgErrors{
			Message:    fmt.Sprintf("ArgsUnknownArgs: %s, for command '%s'", strings.Join(args, ", "), command),
			StatusCode: UnknownArgsErrorCode,
		},
	}
}

func NewFlagArgMissing(flag string, command CommandName) ArgErrors {
	return &FlagArgMissing{
		GenericArgErrors{
			Message:    fmt.Sprintf("ArgsFlagArgsMissing: %s, for command '%s'", flag, command),
			StatusCode: FlagArgMissingErrorCode,
		},
	}
}

func NewFlagCombinationError(args []string, command CommandName) ArgErrors {
	return &FlagCombinationError{
		GenericArgErrors{
			Message:    fmt.Sprintf("ArgsFlagCombinationError: %s, for command '%s'", strings.Join(args, ", "), command),
			StatusCode: FlagCombinationErrorCode,
		},
	}
}
