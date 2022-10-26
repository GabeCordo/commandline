package github.com/GabeCordo/commandline

import (
	"time"
)

type Config struct {
	Version     float32 `json:"version"`
	UserProfile struct {
		FirstName string `json:"first-Name"`
		LastName  string `json:"last-Name"`
		Email     string `json:"email"`
	} `json:"profile"`
}

type TerminateOnCompletion bool

type Command interface {
	Name() string
	Run(cli *CommandLine) TerminateOnCompletion
}

type CommandLine struct {
	Config *Config

	Flags struct {
		Debug  bool
		Create bool
		Delete bool
		Show   bool
	}

	MetaData struct {
		WorkingDirectory string
		TimeCalled       time.Time
	}

	args        []string
	numOfArgs   int
	argsPointer int

	Commands map[string]Command
}
