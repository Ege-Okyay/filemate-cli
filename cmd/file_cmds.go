package cmd

import (
	"fmt"
	"log"

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
// Implementation is pending.
func GetAllFilesCommand(args ...interface{}) {
	// Implementation pending
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
