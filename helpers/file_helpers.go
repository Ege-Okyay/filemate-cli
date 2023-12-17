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

func UploadFile(fileName string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error getting current working directory: %s", err)
	}

	filePath := filepath.Join(dir, fileName)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("files", filepath.Base(filePath))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	res, err := SendHttpRequest("/file/upload", "POST", nil, payload, writer.Boundary(), config.GetUserToken())
	if err != nil {
		return err
	}

	var resFormat map[string]interface{}

	json.NewDecoder(res.Body).Decode(&resFormat)

	if resFormat["error"] != nil {
		fmt.Println(resFormat["error"])
	} else {
		fmt.Println(resFormat["message"])
	}

	return nil
}
