package args

import (
	"fmt"
)

func getCommand(name CommandName) (Command, error) {
	for _, c := range Commands {
		if c.Name == name {
			return c, nil
		}
	}
	return Command{}, CommandNotFoundError{CommandName: string(name)}
}

func HelpPrint(cn CommandName) {
	if cn == GlobalCommand {
		printGlobalHelp()
	} else {
		c, err := getCommand(cn)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s\t\t\t\t%s\n", c.Name, c.Description)
			printFlags(c.Flags)
		}
	}
}

func printGlobalHelp() {
	fmt.Println("Available Commands:")
	fmt.Println()
	for _, c := range Commands {
		if c.Name != GlobalCommand {
			fmt.Printf("%s\t\t\t\t%s\n", c.Name, c.Description)
		}
	}
	fmt.Println("\nGlobal Flags:")
	for _, c := range Commands {
		if c.Name == GlobalCommand && len(c.Flags) > 0 {
			printFlags(c.Flags)
		}
	}
}

func printFlags(flags []Flag) {
	if len(flags) > 0 {
		fmt.Println()
	}
	for _, f := range flags {
		fmt.Printf("%s\t%s\t\t\t%s\n", f.Long, f.Short, f.Description)
	}
}

func ParseArgs(args []string, commandMap map[CommandName]func() error) error {
	if len(args) < 1 {
		HelpPrint(GlobalCommand)
		return nil
	}

	command := CommandName(args[0])

	c, err := getCommand(command)
	if err != nil {
		return CommandNotFoundError{CommandName: string(command)}
	}

	executeFunc, exists := commandMap[command]
	if !exists {
		fmt.Printf("Unknown command: %s\n", command)
		HelpPrint(GlobalCommand)
		return nil
	}

	return parseCommandArgs(args[1:], c, executeFunc)
}

func parseCommandArgs(args []string, c Command, executeFunc func() error) error {
	if len(args) == 0 {
		err := executeFunc()
		return err
	}

	if len(c.Flags) == 0 && len(args) > 0 {
		return UnknownArgsError{Args: args, Command: c.Name}
	}

	// Placeholder if we want to add flag logic
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			HelpPrint(c.Name)
			return nil
		} else {
			return UnknownArgsError{Args: args, Command: c.Name}
		}
	}

	return nil
}
