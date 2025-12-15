package main

import(
	"fmt"
	"os"
	"path/filepath"
)

type UserConfig struct {
	configDir string
}

// Define the permissions for the new directory (0755 is common:
// Owner can read/write/execute, Group/Others can read/execute)
const permissionMode = 0755 


func (c *UserConfig)createOrGetConfig() error {
	c.configDir = filepath.Join(c.configDir, "ComfyFlipper")
	/// See if folder named ConfigFlipper is there
	if err := os.MkdirAll(c.configDir, permissionMode); err != nil {
		return err
	}

	/// At this point, we already confirm the user has the config folder
	fmt.Println("Confirm User has Config Folder")
	return nil
}

func main() {
	var config UserConfig
	dir, err := os.UserConfigDir()

	if err != nil {
		fmt.Println("Error Finding Config Directory")
		return
	}

	config.configDir = dir

	err = config.createOrGetConfig()
	if err != nil {
		fmt.Println("Error Creating Config Directory")
		return
	}
}
