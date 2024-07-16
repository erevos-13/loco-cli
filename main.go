package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"voutsaridis.com/loco-translate-cli/api"
	"voutsaridis.com/loco-translate-cli/filemanage"
)

type Translation struct{}

// WriteInFile writes data to a file, creating the file if it does not exist.
func WriteInFile(filePath, data string) {
	// Extract the directory part from the filePath
	dir := filepath.Dir(filePath)

	// Create the directory if it does not exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// If the file does not exist, create it
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("Failed to create file: %s", err)
		}
		defer file.Close()
	} else {
		// If the file exists, open it in append mode
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
		}
		defer file.Close()

		// Write data to the file
		if _, err := file.WriteString(data); err != nil {
			log.Fatalf("Failed to write to file: %s", err)
		}
	}
}

func main() {
	root, _ := os.Getwd()
	filePath := root + "/data/%s.json"
	writeFilePath := root + "/assets/locale/%s.json"

	stringJson, err := filemanage.ReadTextFile(fmt.Sprintf(filePath, "en"))
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	sendFile, errorSendFile := api.PostTranslation(stringJson)
	if errorSendFile != nil {
		panic(errorSendFile)
	}
	fmt.Println(sendFile)

	response, err := api.GetTranslationByLocal("en")
	if err != nil {
		panic(err)
	}
	urlPath := fmt.Sprintf(writeFilePath, "en")
	WriteInFile(urlPath, response)

}
