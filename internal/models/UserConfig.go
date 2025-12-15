package models


import(
	"fmt"
	"os"
	"path/filepath"
)


type UserConfig struct {
	ConfigDir string
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
	fmt.Println("Confirm User has Config Folder")
	return nil
}

