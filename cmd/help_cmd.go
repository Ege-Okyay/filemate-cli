package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/Ege-Okyay/filemate-cli/helpers"
)

// HelpCmd handles the "help" command, providing usage and description information for a specific command,
// or listing all available commands if no command is specified.
func HelpCmd(args ...interface{}) {
	// Check if command name is provided as an argument
	if len(args) > 0 {
		if argSlice, ok := args[0].([]interface{}); ok {
			if len(argSlice) > 0 {
				// Extract command name from the arguments
				cmdName := argSlice[0].(string)

				// Find the command by name
				cmd, err := helpers.FindCommandByName(cmdName)
				if err != nil {
					log.Fatal("Error while finding command: ", err)
				}

				// Print usage and description of the specified command
				fmt.Printf("Usage: %s\n%s", cmd.Usage, cmd.Desc)

				return
			} else {
				log.Fatal("Empty slice of strings")
				return
			}
		}
	}

	// If no command name is provided, list all available commands
	commands := helpers.GetAllCommands()

	// Define a template for displaying usage information
	tmpl := `Filemate is a tool for uploading, downloading and sharing files from your terminal.

Usage:

	filemate <command> [arguments]

The commands are:
{{range .}}
	{{.Name}}{{end}}

Use "filemate help <command>" for more information about a command.
		`

	// Create a template and execute it to display usage information
	t := template.Must(template.New("usage").Parse(tmpl))
	err := t.Execute(os.Stdout, commands)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}
}
