package models

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)


type UserConfig struct {
	ConfigDir string
	CurrentDirectory CurrentDirectory
}

// Define the permissions for the new directory (0755 is common:
// Owner can read/write/execute, Group/Others can read/execute)
const permissionMode = 0755 


func (c *UserConfig)CreateOrGetConfig() error {
	c.ConfigDir = filepath.Join(c.ConfigDir, "ConfigFlipper")
	fmt.Println(c.ConfigDir)
	/// See if folder named ConfigFlipper is there
	if err := os.MkdirAll(c.ConfigDir, permissionMode); err != nil {
		return err
	}

	/// At this point, we already confirm the user has the config folder
	return nil
}

func (c UserConfig)OpenFile() error {
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
