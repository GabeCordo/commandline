package commandline

import (
	"fmt"
)

// HELP COMMAND START

type HelpCommand struct {
}

func (helpCommand HelpCommand) Run(cli *CommandLine) TerminateOnCompletion {
	fmt.Println("Default commands:")
	fmt.Println("help\tView helpful information about the etl service")

	if len(cli.commands) > 0 {
		fmt.Println("Core commands:")

		for _, command := range cli.commands {
			fmt.Println(command.Name())
		}
	}

	return Terminate
}

// HELP SCOPE COMMAND

type HelpScopeCommand struct {
}

func (helpScopeCommand HelpScopeCommand) Run(cli *CommandLine) TerminateOnCompletion {
	if len(cli.args) == 1 {
		fmt.Println("you cannot use '?' without specifying a command")
	}

	if (len(cli.args) == 2) && (cli.args[1] == "?") {
		if commandWrapper, found := cli.commands[cli.args[0]]; found {
			fmt.Println(commandWrapper.description)
		} else {
			fmt.Println(cli.args[0] + " is not a command")
		}
	} else if (len(cli.args) == 3) && (cli.args[2] == "?") {
		if commandWrapper, found := cli.commands[cli.args[0]]; found {
			if variant, found := commandWrapper.variants[cli.args[1]]; found {
				fmt.Println(variant)
			} else {
				fmt.Printf("%s [%s] either does not exist or does not have a description")
			}
		}
	} else {
		fmt.Println("invalid format, you must use '[command] ?' or '[command] [flag] ?'")
	}

	return Terminate
}
