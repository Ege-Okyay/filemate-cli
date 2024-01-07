package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
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

	files, err := helpers.GetFiles()
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
func DownloadFileCommand(args ...interface{}) {
	// Check if the user is logged in
	if config.GetUserToken() == "" {
		fmt.Println("Login first")
		return
	}

	fileName := args[0].([]interface{})[0].(string)
	files, err := helpers.GetFiles()
	if err != nil {
		log.Fatal(err)
	}

	var fileID string
	for _, file := range files {
		if file.FileName == fileName {
			fileID = file.ID
			break
		}
	}

	route := fmt.Sprintf("/file/download?fileId=%s", fileID)
	res, err := helpers.SendHttpRequest(route, "GET", nil, nil, "", config.GetUserToken())
	if err != nil {
		log.Fatal(err)
	}

	var resFormat map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resFormat)

	if resFormat["error"] != nil {
		fmt.Println(resFormat["error"])
		return
	}

	contentDisposition := res.Header.Get("Content-Disposition")
	if contentDisposition == "" {
		fmt.Println("Error: Content-Disposition header not found in the response.")
		return
	}

	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		fmt.Println("Error parsing Content-Disposition header: ", err)
		return
	}

	fileNameFromHeader := params["filename"]

	file, err := os.Create(fileNameFromHeader)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	fmt.Println("---DEBUG--- RESPONSE BODY : \n", res.Body)

	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Println("Error copying response body to file: ", err)
		return
	}

	fmt.Println("File downloaded successfully: ", fileNameFromHeader)
}

// DeleteFileCommand handles the "delete-file" command, allowing users to delete a specified file from the server.
// Implementation is pending.
func DeleteFileCommand(args ...interface{}) {
	// Implementation pending
}
