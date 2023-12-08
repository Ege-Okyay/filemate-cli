package cli

import (
	"fmt"
	"os"

	"github.com/Ege-Okyay/filemate-cli/cmd"
	"github.com/Ege-Okyay/filemate-cli/helpers"
)

func registerCommands() {
	helpers.RegisterCommand("hello", cmd.HelloCmd, "Says Hello", "hello [name]")
	helpers.RegisterCommand("--help", cmd.HelpCmd, "Shows Help", "help")
}

func CliProgram() error {
	registerCommands()

	args := os.Args[1:]
	if len(args) == 0 {
		return fmt.Errorf("Please provide arg(s)")
	}

	cmdArg := args[0]

	cmd, err := helpers.FindCommandByName(cmdArg)
	if err != nil {
		return err
	}

	argsForFunction := helpers.ExtractArguments(args)
	helpers.CallCommand(cmd, argsForFunction)

	return nil
}
