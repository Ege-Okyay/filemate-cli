package cmd

import (
	"fmt"
	"os"
)

func VersionCommand(args ...interface{}) {
	currentVersion := os.Getenv("VERSION")

	fmt.Printf("filemate version %s\n", currentVersion)
}
