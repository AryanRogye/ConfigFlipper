package models

import (
	"os"
)

type CurrentDirectory struct {
	data []currentDirectoryData
}
func (c *CurrentDirectory) addData(name string, isDir bool) {
	if isDir {
		c.data = append(c.data, Folder{name})
	} else {
		c.data = append(c.data, File{name})
	}
}

func GetCurrentDirectoryData() (CurrentDirectory, error) {
	currentDirectory, err := os.Getwd()

	var cd CurrentDirectory;
	if err != nil {
		return cd, err
	}


	entries, err := os.ReadDir(currentDirectory)

	if err != nil {
		return  cd, err
	}

	for _, entry := range(entries) {
		cd.addData(entry.Name(), entry.IsDir())
	}

	return cd, nil
}

type currentDirectoryData interface {
	currentDirNode()
}

type File struct {
	name string
}
func (File) currentDirNode() {}


type Folder struct {
	name string
}
func (Folder) currentDirNode() {}
