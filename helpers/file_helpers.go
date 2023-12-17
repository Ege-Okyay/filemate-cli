package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/Ege-Okyay/filemate-cli/config"
)

// UploadFile uploads the specified file to the server.
// It takes the file name as an argument and returns an error if any occurs during the process.
func UploadFile(fileName string) error {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error getting current working directory: %s", err)
	}

	// Construct the full file path
	filePath := filepath.Join(dir, fileName)

	// Create a buffer to store the payload
	payload := &bytes.Buffer{}

	// Create a new multipart writer for constructing the HTTP request
	writer := multipart.NewWriter(payload)

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	// Create a form file part in the multipart writer
	part, err := writer.CreateFormFile("files", filepath.Base(filePath))
	if err != nil {
		return err
	}

	// Copy the file content to the form file part
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return err
	}

	// Send the HTTP request to upload the file
	res, err := SendHttpRequest("/file/upload", "POST", nil, payload, writer.Boundary(), config.GetUserToken())
	if err != nil {
		return err
	}

	// Decode the response JSON
	var resFormat map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resFormat)

	// Check if the sign-up was successful
	if resFormat["error"] != nil {
		fmt.Println(resFormat["error"])
	} else {
		fmt.Println(resFormat["message"])
	}

	return nil
}
