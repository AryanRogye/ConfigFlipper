package models

import (
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
}

// Define the permissions for the new directory (0755 is common:
// Owner can read/write/execute, Group/Others can read/execute)
const permissionMode = 0755

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

// Config's of the same name can be stored as:
// Config(1)
// Config(2)
// Config(3)
// ...
// We want to return the number of configs of the same name
func (c UserConfig) GetNConfigCount(config_name string) int {
	entries, err := os.ReadDir(c.ConfigDir)
	if err != nil {
		return 0
	}

	count := 0
	for _, entry := range entries {

		name := entry.Name()

		if name == config_name {
			count++
			continue
		}

		// match "Config(n)"
		open := strings.LastIndex(name, "(")
		close := strings.LastIndex(name, ")")
		if open == -1 || close != len(name)-1 || close < open {
			continue
		}
		base := name[:open]
		if base == config_name {
			count++
		}
	}
	return count
}

func (c UserConfig) CreateConfig(config_name string) error {
	path := filepath.Join(c.ConfigDir, config_name)
	return os.Mkdir(path, permissionMode)
}

// Returns true if the config exists
func (c UserConfig) DoesConfigExist(config_name string) bool {
	path := filepath.Join(c.ConfigDir, config_name)
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Opens the config folder
func (c UserConfig) OpenConfigFolder() error {
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
