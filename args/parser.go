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

func ParseArgs(args []string) {
	if len(args) < 1 {
		HelpPrint(GlobalCommand)
		return
	}

	command := args[0]
	switch command {
	case string(ClusterCommand):
		parseCommandArgs(args[1:], ClusterCommand)
	case string(DeployCommand):
		parseCommandArgs(args[1:], DeployCommand)
	case string(StatusCommand):
		parseCommandArgs(args[1:], StatusCommand)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		HelpPrint(GlobalCommand)
	}
}

func parseCommandArgs(args []string, cmd CommandName) {
	if len(args) == 0 {
		fmt.Println("Executing", cmd, "command...")
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
