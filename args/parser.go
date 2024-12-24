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
	for _, f := range flags {
		fmt.Printf("\t%s\t%s\t%s\n", f.Long, f.Short, f.Description)
	}
}

func ParseArgs(args []string, commandMap map[CommandName]func()) {
	if len(args) < 1 {
		HelpPrint(GlobalCommand)
		return
	}

	command := CommandName(args[0])
	executeFunc, exists := commandMap[command]
	if !exists {
		fmt.Printf("Unknown command: %s\n", command)
		HelpPrint(GlobalCommand)
		return
	}

	parseCommandArgs(args[1:], command, executeFunc)
}

func parseCommandArgs(args []string, cmd CommandName, executeFunc func()) {
	if len(args) == 0 {
		executeFunc()
		return
	}

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			HelpPrint(cmd)
			return
		} else {
			fmt.Println(InvalidFlagError{FlagName: arg, Command: cmd})
		}
	}
}
