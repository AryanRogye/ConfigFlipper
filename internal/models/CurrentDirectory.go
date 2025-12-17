package models

import (
	"os"
)

type CurrentDirectory struct {
	Data []CurrentDirectoryData
}
func (c *CurrentDirectory) addData(name string, isDir bool) {
	if isDir {
		c.Data = append(c.Data, Folder{name})
	} else {
		c.Data = append(c.Data, File{name})
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

type CurrentDirectoryData interface {
	Name() string
}

type File struct {
	name string
}
func (f File) Name() string {
	return f.name
}

type Folder struct {
	name string
}
func (f Folder) Name() string {
	return f.name
}
