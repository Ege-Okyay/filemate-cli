package cmd

import (
	"fmt"
	"log"

	"github.com/Ege-Okyay/filemate-cli/config"
	"github.com/Ege-Okyay/filemate-cli/helpers"
)

func UploadFileCommand(args ...interface{}) {
	if config.GetUserToken() == "" {
		fmt.Println("Login first")
		return
	}

	fileName := args[0].([]interface{})[0].(string)

	err := helpers.UploadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllFilesCommand(args ...interface{}) {

}

func DownloadFileCommand(args ...interface{}) {

}

func DeleteFileCommand(args ...interface{}) {

}
