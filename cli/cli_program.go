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
	helpers.RegisterCommand("hello", cmd.HelloCmd, "DEBUG COMMAND Says Hello", "filemate hello [name]", "Debug", 1, false)
	helpers.RegisterCommand("health-check", cmd.HealthCheckCmd, "DEBUG COMMAND Sends health check request to the api", "filemate health-check", "Debug", 0, false)
	helpers.RegisterCommand("post-check", cmd.PostCheckCmd, "DEBUG COMMAND Sends post check request to the api", "filemate post-check [arguments]", "Debug", 1, false)

	// INFO COMMANDS
	helpers.RegisterCommand("help", cmd.HelpCmd, "Shows Help", "filemate help", "Info", 0, true)
	helpers.RegisterCommand("version", cmd.VersionCommand, "Shows version", "filemate version", "Info", 0, false)

	// AUTH COMMANDS
	helpers.RegisterCommand("sign-up", cmd.SignUpCommand, "Creates an account with given username, password and email", "filemate sign-up [username] [email] [password]", "Authentication", 3, false)
	helpers.RegisterCommand("login", cmd.LoginCommand, "Logs in the user with the entered username or email and password", "filemate login [username or email] [password]", "Authentication", 2, false)
	helpers.RegisterCommand("logout", cmd.LogoutCommand, "Logs out the user", "filemate logout", "Authentication", 0, false)
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
