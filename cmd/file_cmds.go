package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/Ege-Okyay/filemate-cli/config"
	"github.com/Ege-Okyay/filemate-cli/helpers"
)

// UploadFileCommand handles the "upload-file" command, allowing users to upload a specified file to the server.
func UploadFileCommand(args ...interface{}) {
	// Check if the user is logged in
	if config.GetUserToken() == "" {
		fmt.Println("Login first")
		return
	}

	// Extract the file name from the arguments
	fileName := args[0].([]interface{})[0].(string)

	// Call the helper function to upload the file
	err := helpers.UploadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllFilesCommand handles the "get-all-files" command, allowing users to retrieve a list of all files from the server.
// It checks if the user is logged in, sends a request to the server to get the list of files, and displays the file details.
// Usage: filemate get-all-files
// Permissions: Requires user to be logged in.
func GetAllFilesCommand(args ...interface{}) {
	// Check if the user is logged in
	if config.GetUserToken() == "" {
		fmt.Println("Login first")
		return
	}

	// Send a GET request to retrieve the list of files from the server
	res, err := helpers.SendHttpRequest("/file/files", "GET", nil, nil, "", config.GetUserToken())
	if err != nil {
		log.Fatal(err)
	}

	// Decode the response body into a map
	var resFormat map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resFormat)

	// Check for errors in the response
	if resFormat["error"] != nil {
		fmt.Println(resFormat["error"])
		return
	}

	// Extract the raw files data from the response
	filesRaw, ok := resFormat["files"].([]interface{})
	if !ok {
		log.Fatal("Invalid or missing 'files' key in the response.")
	}

	// Unmarshal the raw files data into a slice of File structs
	files, err := helpers.UnmarshalFileEntires(filesRaw)
	if err != nil {
		log.Fatal(err)
	}

	// Define a template for displaying file details
	tmpl := `
Your files are:
{{range .}}
	File Name: {{.FileName}} - File Size: {{FormatFileSize .FileSize}}{{end}}

Use "filemate file-details [file name]" for more information about a file.
`

	// Create a template and execute it to display file details
	t := template.Must(template.New("usage").Funcs(template.FuncMap{"FormatFileSize": helpers.FormatFileSize}).Parse(tmpl))
	err = t.Execute(os.Stdout, files)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}
}

// DownloadFileCommand handles the "download-file" command, allowing users to download a specified file from the server.
// Implementation is pending.
func DownloadFileCommand(args ...interface{}) {
	// Implementation pending
}

// DeleteFileCommand handles the "delete-file" command, allowing users to delete a specified file from the server.
// Implementation is pending.
func DeleteFileCommand(args ...interface{}) {
	// Implementation pending
}
