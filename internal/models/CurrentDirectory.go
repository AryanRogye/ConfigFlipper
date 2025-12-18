package models

import (
	"os"
	"path/filepath"
)

type CurrentDirectory struct {
	Data []CurrentDirectoryData
}

func (c *CurrentDirectory) addData(sourceDir string, path string, name string, isDir bool) {
	if isDir {
		c.Data = append(c.Data, Folder{name, path, sourceDir})
	} else {
		c.Data = append(c.Data, File{name, path, sourceDir})
	}
}

func GetCurrentDirectoryData() (CurrentDirectory, error) {
	currentDirectory, err := os.Getwd()

	var cd CurrentDirectory
	if err != nil {
		return cd, err
	}

	entries, err := os.ReadDir(currentDirectory)

	if err != nil {
		return cd, err
	}

	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(currentDirectory, name)
		cd.addData(currentDirectory, path, name, entry.IsDir())
	}

	return cd, nil
}

type CurrentDirectoryData interface {
	Name() string
	FullPath() string
	SourceDir() string
}

type File struct {
	name      string
	fullPath  string
	sourceDir string
}

func (f File) Name() string      { return f.name }
func (f File) FullPath() string  { return f.fullPath }
func (f File) SourceDir() string { return f.sourceDir }

type Folder struct {
	name      string
	fullPath  string
	sourceDir string
}

func (f Folder) Name() string      { return f.name }
func (f Folder) FullPath() string  { return f.fullPath }
func (f Folder) SourceDir() string { return f.sourceDir }
