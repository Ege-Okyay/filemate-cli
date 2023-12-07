package main

import (
	"fmt"

	"github.com/Ege-Okyay/filemate-cli/cli"
)

func main() {
	err := cli.CliProgram()
	if err != nil {
		fmt.Printf("CLI ERROR: %s", err)
		return
	}
}
