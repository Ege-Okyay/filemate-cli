package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/Ege-Okyay/filemate-cli/helpers"
)

func HelpCmd(args []interface{}) {
	if len(args) > 0 {
		cmdName := args[0].(string)

		cmd, err := helpers.FindCommandByName(cmdName)
		if err != nil {
			log.Fatal("Error while finding command: ", err)
		}

		fmt.Printf("Usage: %s\n%s", cmd.Usage, cmd.Desc)

		return
	}

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
}
