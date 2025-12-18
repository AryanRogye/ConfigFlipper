package models

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type UserConfig struct {
	ConfigDir        string
	CurrentDirectory CurrentDirectory
	Data             CurrentDirectoryData
}

// Define the permissions for the new directory (0755 is common:
// Owner can read/write/execute, Group/Others can read/execute)
const permissionMode = 0755

// CreateOrGetConfig
// Function will create the config directory if it doesn't exist
// Either or, ConfigDir will be set
func (c *UserConfig) CreateOrGetConfig() error {
	c.ConfigDir = filepath.Join(c.ConfigDir, "ConfigFlipper")
	fmt.Println(c.ConfigDir)
	/// See if folder named ConfigFlipper is there
	if err := os.MkdirAll(c.ConfigDir, permissionMode); err != nil {
		return err
	}

	/// At this point, we already confirm the user has the config folder
	return nil
}

// GetNConfigCount
// Config's of the same name can be stored as:
// Config(1)
// Config(2)
// Config(3)
// ...
// We want to return the number of configs of the same name
func (c *UserConfig) GetNConfigCount(configName string) int {
	entries, err := os.ReadDir(c.ConfigDir)
	if err != nil {
		return 0
	}

	count := 0
	for _, entry := range entries {

		name := entry.Name()

		if name == configName {
			count++
			continue
		}

		// match "Config(n)"
		open := strings.LastIndex(name, "(")
		last := strings.LastIndex(name, ")")
		if open == -1 || last != len(name)-1 || last < open {
			continue
		}
		base := name[:open]
		if base == configName {
			count++
		}
	}
	return count
}

func (c *UserConfig) CreateConfig(configName string) error {
	if c.Data == nil {
		return errors.New("no data found")
	}
	path := filepath.Join(c.ConfigDir, configName)
	err := os.Mkdir(path, permissionMode)

	if err == nil {
		/// Create SymLink for Data

		// First we get the original source path
		sourcePath := c.Data.SourceDir()

		// Get Source File System
		srcFS := os.DirFS(sourcePath)

		/// Copy
		err = os.CopyFS(path, srcFS)
		if err != nil {
			return err
		}
	}
	return err
}

// OpenConfigFolder
// Opens the config folder
func (c *UserConfig) OpenConfigFolder() error {
	path := c.ConfigDir
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", path)
	case "windows":
		cmd = exec.Command("explorer", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
