package commandline

type Command interface {
	Run(cli *CommandLine) TerminateOnCompletion
}

type CommandWrapper struct {
	command     Command // IMMUTABLE
	identifier  string  // IMMUTABLE
	description string
	category    string
	variants    map[string]string
}

func NewCommandWrapper(identifier string, command Command) *CommandWrapper {
	if command == nil {
		panic("nil memory address cannot be passed to the CommandWrapper")
	}

	commandWrapper := new(CommandWrapper)
	commandWrapper.command = command
	commandWrapper.identifier = identifier
	commandWrapper.variants = make(map[string]string)

	return commandWrapper
}

func (commandWrapper CommandWrapper) Name() string {
	return commandWrapper.identifier
}

func (commandWrapper CommandWrapper) Command() Command {
	return commandWrapper.command
}

func (commandWrapper *CommandWrapper) Help(variant ...string) string {
	// we can take 0 to 1 arguments in this function
	if !(len(variant) == 0 || len(variant) == 1) {
		panic("CommandWrapper.Help takes 0 to 1 arguments, too many were passed into it")
	}

	// if no variant type was passed to this function, output the default command description
	if len(variant) == 1 {
		return commandWrapper.description
	} else {
		return commandWrapper.variants[variant[0]]
	}
}

func (commandWrapper *CommandWrapper) SetDescription(args ...string) *CommandWrapper {
	if !((len(args) == 0) || (len(args) == 1)) {
		panic("CommandWrapper.Description takes 0 to 1 arguments, too many were passed into it")
	}

	if len(args) == 1 {
		commandWrapper.description = args[0]
	} else if len(args) == 2 {
		commandWrapper.variants[args[0]] = args[1]
	}

	return commandWrapper
}

func (commandWrapper CommandWrapper) Description(args ...string) string {
	if !((len(args) == 0) || (len(args) == 0)) {
		panic("CommandWrapper.Description takes 0 to 1 arguments, too many were passed into it")
	}

	if len(args) == 1 {
		return commandWrapper.description
	} else {
		return commandWrapper.variants[args[0]]
	}
}

func (commandWrapper *CommandWrapper) SetCategory(category string) *CommandWrapper {
	commandWrapper.category = category
	return commandWrapper
}

func (commandWrapper CommandWrapper) Category() string {
	return commandWrapper.category
}
