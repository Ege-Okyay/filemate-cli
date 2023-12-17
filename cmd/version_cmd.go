package cmd

import (
	"fmt"
	"os"
)

// VersionCommand handles the "version" command, displaying the current version of the Filemate CLI.
func VersionCommand(args ...interface{}) {
	// Retrieve the current version from the environment variable
	currentVersion := os.Getenv("VERSION")

	// Print the Filemate CLI version
	fmt.Printf("filemate version %s\n", currentVersion)
}
