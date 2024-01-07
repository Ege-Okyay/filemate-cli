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
	"github.com/Ege-Okyay/filemate-cli/structs"
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

func GetFiles() ([]structs.File, error) {
	// Send a GET request to retrieve the list of files from the server
	res, err := SendHttpRequest("/file/files", "GET", nil, nil, "", config.GetUserToken())
	if err != nil {
		return nil, err
	}

	// Decode the response body into a map
	var resFormat map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resFormat)

	// Check for errors in the response
	if resFormat["error"] != nil {
		return nil, fmt.Errorf(resFormat["error"].(string))
	}

	// Extract the raw files data from the response
	filesRaw, ok := resFormat["files"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Invalid or missing 'files' key in the response.")
	}

	// Unmarshal the raw files data into a slice of File structs
	files, err := UnmarshalFileEntires(filesRaw)
	if err != nil {
		return nil, err
	}

	// Return the files
	return files, nil
}

func UnmarshalFileEntires(filesRaw []interface{}) ([]structs.File, error) {
	var files []structs.File

	for _, fileRaw := range filesRaw {
		fileBytes, err := json.Marshal(fileRaw)
		if err != nil {
			return nil, err
		}

		var file structs.File
		err = json.Unmarshal(fileBytes, &file)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

// FormatFileSize formats the given file size in bytes to a human-readable format.
func FormatFileSize(size int64) string {
	const (
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	var unit string
	var value float64

	switch {
	case size < KB:
		unit = "B"
		value = float64(size)
	case size < MB:
		unit = "KB"
		value = float64(size) / KB
	case size < GB:
		unit = "MB"
		value = float64(size) / MB
	case size < TB:
		unit = "GB"
		value = float64(size) / GB
	case size < PB:
		unit = "TB"
		value = float64(size) / TB
	case size < EB:
		unit = "PB"
		value = float64(size) / PB
	default:
		unit = "EB"
		value = float64(size) / EB
	}

	return fmt.Sprintf("%.2f %s", value, unit)
}
