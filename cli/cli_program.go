package cli

import (
	"log"
	"os"
	"text/template"

	"github.com/Ege-Okyay/filemate-cli/cmd"
	"github.com/Ege-Okyay/filemate-cli/helpers"
)

// registerCommands registers all the available commands for the Filemate CLI.
func registerCommands() {
	// INFO COMMANDS
	helpers.RegisterCommand("help", cmd.HelpCmd, "Shows Help", "filemate help", "Info", 0, true)
	helpers.RegisterCommand("version", cmd.VersionCommand, "Shows version", "filemate version", "Info", 0, false)

	// AUTH COMMANDS
	helpers.RegisterCommand("sign-up", cmd.SignUpCommand, "Creates an account with given username, password and email", "filemate sign-up [username] [email] [password]", "Authentication", 3, false)
	helpers.RegisterCommand("login", cmd.LoginCommand, "Logs in the user with the entered username or email and password", "filemate login [username or email] [password]", "Authentication", 2, false)
	helpers.RegisterCommand("logout", cmd.LogoutCommand, "Logs out the user", "filemate logout", "Authentication", 0, false)

	// FILE COMMANDS
	helpers.RegisterCommand("upload-file", cmd.UploadFileCommand, "Uploads the specified file to the server", "filemate upload-file [file name]", "File", 1, false)
}

// CliProgram is the main entry point for the Filemate CLI program.
// It registers commands, processes command-line arguments, and executes the requested command.
func CliProgram() error {
	// Register all available commands
	registerCommands()

	// Get command-line arguments excluding the program name
	args := os.Args[1:]

	// If no command is provided, display general usage information
	if len(args) == 0 {
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

		return nil
	}

	// Extract the command provided as the first argument
	cmdArg := args[0]

	// Find the command by name
	cmd, err := helpers.FindCommandByName(cmdArg)
	if err != nil {
		return err
	}

	// Extract arguments for the command
	argsForFunction := helpers.ExtractArguments(args)

	// Call the selected command with the provided arguments
	helpers.CallCommand(cmd, argsForFunction)

	return nil
}
