package github.com/GabeCordo/commandline

import (
	"errors"
	"io/ioutil"
	"os"
)

type PathType uint8

const (
	FilePath      PathType = 0
	DirectoryPath          = 1
)

const (
	DefaultFilePermissions os.FileMode = 0755
)

type Path struct {
	value string
	ptype PathType
}

func EmptyPath() Path {
	return Path{value: "", ptype: DirectoryPath}
}

func (path Path) IsFile() bool {
	return path.ptype == FilePath
}

func (path Path) IsDirectory() bool {
	return path.ptype == DirectoryPath
}

func (path Path) ToString() string {
	return path.value
}

func (path Path) BackDir() Path {
	if path.IsFile() {
		panic("you cannot chain a directory path to a file")
	}

	return Path{path.value + "/../", DirectoryPath}
}

func (path Path) Dir(directory string) Path {
	if path.IsFile() {
		panic("you cannot chain a directory to a file")
	}

	return Path{path.value + "/" + directory + "/", DirectoryPath}
}

func (path Path) File(file string) Path {
	if path.IsFile() {
		panic("you cannot chain a file to a file")
	}

	return Path{path.value + file, FilePath}
}

func (path Path) Exists() bool {
	_, err := os.Stat(path.value)
	return os.IsExist(err)
}

func (path Path) DoesNotExist() bool {
	return !path.Exists()
}

func (path Path) Read() ([]byte, error) {
	if path.DoesNotExist() {
		return []byte{}, errors.New(path.value + " does not exist")
	}

	if path.IsDirectory() {
		return []byte{}, errors.New(path.value + " is not a file")
	}

	bytes, _ := ioutil.ReadFile(path.value)
	return bytes, nil
}

func (path Path) Write(bytes []byte) error {
	if path.DoesNotExist() {
		return errors.New(path.value + " does not exist")
	}

	if path.IsDirectory() {
		return errors.New(path.value + " is not a file")
	}

	return ioutil.WriteFile(path.value, bytes, DefaultFilePermissions)
}

func (path Path) Remove() error {
	if path.DoesNotExist() {
		return errors.New(path.value + " does not exist")
	}

	return os.RemoveAll(path.value)
}

func (path Path) MkDir() error {
	if path.Exists() {
		return errors.New(path.value + " already exists")
	}

	if path.IsFile() {
		return errors.New("cannot create a directory with the path for a file")
	}

	return os.Mkdir(path.value, DefaultFilePermissions)
}
