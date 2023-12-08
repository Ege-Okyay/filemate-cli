package cmd

import (
	"fmt"

	"github.com/Ege-Okyay/filemate-cli/helpers"
)

func HelpCmd(args []interface{}) {
	allCommands := helpers.GetAllCommands()

	fmt.Println("All Commands\n------------")
	fmt.Println("usage: git [-v | --version] [-h | --help] [-C <path>] [-c <name>=<value>]\n[--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]\n[-p | --paginate | -P | --no-pager] [--no-replace-objects] [--bare]\n[--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]\n[--config-env=<name>=<envvar>] <command> [<args>]")
	for key := range allCommands {
		fmt.Println(key)
	}
}
