package cli

import (
	"log"
	"os"
	"text/template"

	"github.com/Ege-Okyay/filemate-cli/cmd"
	"github.com/Ege-Okyay/filemate-cli/helpers"
)

func registerCommands() {
	// DEBUG COMMANDS
	helpers.RegisterCommand("hello", cmd.HelloCmd, "DEBUG COMMAND Says Hello", "filemate hello [args]", true)
	helpers.RegisterCommand("health-check", cmd.HealthCheckCmd, "DEBUG COMMAND Sends health check request to the api", "filemate health-check", false)

	helpers.RegisterCommand("help", cmd.HelpCmd, "Shows Help", "filemate help", false)
	helpers.RegisterCommand("version", cmd.VersionCommand, "Shows version", "filemate version", false)
}

func CliProgram() error {
	registerCommands()

	args := os.Args[1:]

	if len(args) == 0 {
		commands := helpers.GetAllCommands()

		tmpl := `Filemate is a tool for uploading, downloading and sharing files from your terminal.

Usage:

	filemate <command> [arguments]			

The commands are:
{{range .}}
	{{.Name}}{{end}}

Use "filemate help <command>" for more information about a command.
		`

		t := template.Must(template.New("usage").Parse(tmpl))
		err := t.Execute(os.Stdout, commands)
		if err != nil {
			log.Fatal("Error executing template: ", err)
		}

		return nil
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
