package commandline

import (
	"fmt"
)

// DEBUG COMMAND START

type DebugCommand struct {
	name string
}

func (dc DebugCommand) Name() string {
	return dc.name
}

func (dc DebugCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.Flags.Debug = true

	return false // do not terminate
}

// CREATE COMMAND START

type CreateCommand struct {
	name string
}

func (cc CreateCommand) Name() string {
	return cc.name
}

func (cc CreateCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.Flags.Create = true

	return false // do not terminate
}

// DELETE COMMAND START

type DeleteCommand struct {
	name string
}

func (dc DeleteCommand) Name() string {
	return dc.name
}

func (dc DeleteCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.Flags.Delete = true

	return false // do not terminate
}

// SHOW COMMAND START

type ShowCommand struct {
	name string
}

func (sc ShowCommand) Name() string {
	return sc.name
}

func (sc ShowCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.Flags.Show = true

	return false
}

// INSTALL COMMAND START

type InstallCommand struct {
	name string
}

func (ic InstallCommand) Name() string {
	return ic.name
}

func (ic InstallCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.Flags.Install = true

	return false
}

// INSTALL COMMAND END

// HELP COMMAND START

type HelpCommand struct {
	name string
}

func (helpCommand HelpCommand) Name() string {
	return helpCommand.name
}

func (helpCommand HelpCommand) Run(cli *CommandLine) TerminateOnCompletion {
	fmt.Println("Default Commands:")
	fmt.Println("help\tView helpful information about the etl service")
	
	if len(cli.Commands) > 0 {
		fmt.Println("Core Commands:")

		for _, command := range cli.Commands {
			fmt.Println(command.Name())
		}
	}

	return true
}