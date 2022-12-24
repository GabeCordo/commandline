package commandline

import "github.com/GabeCordo/toolchain/strings"

type Flag uint32

const (
	Debug Flag = iota
	Create
	Delete
	Show
	Install
	Add
	Update
	Revoke
	NotAFlag
)

const (
	numOfFlags uint8 = 9
)

var strToFlagArr = [...]string{"debug", "create", "delete", "show", "install", "add", "update", "revoke"}

func (flag Flag) ToString() string {
	if flag == NotAFlag {
		return strings.EmptyString
	}
	return strToFlagArr[flag]
}

func FlagFromString(strFlag string) Flag {
	if strFlag == "debug" {
		return Debug
	} else if strFlag == "create" {
		return Create
	} else if strFlag == "delete" {
		return Delete
	} else if strFlag == "show" {
		return Show
	} else if strFlag == "install" {
		return Install
	} else if strFlag == "add" {
		return Add
	} else if strFlag == "update" {
		return Update
	} else if strFlag == "revoke" {
		return Revoke
	} else {
		return NotAFlag
	}
}
