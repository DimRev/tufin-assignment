package args

import (
	"fmt"
	"strings"
)

type CommandNotFoundError struct {
	CommandName string
}

func (e CommandNotFoundError) Error() string {
	return fmt.Sprintf("Command %s not found", e.CommandName)
}

type InvalidFlagError struct {
	FlagName string
	Command  CommandName
}

func (e InvalidFlagError) Error() string {
	return fmt.Sprintf("Invalid flag '%s' for command '%s'", e.FlagName, e.Command)
}

type UnknownArgsError struct {
	Args    []string
	Command CommandName
}

func (e UnknownArgsError) Error() string {
	return fmt.Sprintf("Unknown arguments '%s' for command '%s'", strings.Join(e.Args, ", "), e.Command)
}
