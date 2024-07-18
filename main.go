package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/erevos-13/loco-cli/api"
	"github.com/erevos-13/loco-cli/filemanage"
	"github.com/erevos-13/loco-cli/locales"
)

type Translation struct{}

// WriteInFile writes data to a file, creating the file if it does not exist.
func WriteInFile(filePath, data string) {
	// Extract the directory part from the filePath
	dir := filepath.Dir(filePath)

	// Create the directory if it does not exist
	if err := os.MkdirAll(dir, 0777); err != nil {
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
		// Write data to the file
		if _, err := file.WriteString(data); err != nil {
			log.Fatalf("Failed to write to file: %s", err)
		}
	} else {
		// If the file exists, open it in append mode
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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
	argsWithProg := os.Args[1:]
	fmt.Println("-------------------")
	fmt.Printf("Arguments Token: %s\n", argsWithProg[0])
	fmt.Println("-------------------")
	fmt.Printf("Arguments Path to get the json: %s\n", argsWithProg[1])
	fmt.Println("-------------------")
	fmt.Printf("Arguments locale: %s\n", argsWithProg[2])
	fmt.Println("-------------------")
	fmt.Printf("Arguments path to store the json: %s \n", argsWithProg[3])
	fmt.Println("-------------------")

	stringJson, err := filemanage.ReadTextFile(root + string(argsWithProg[1]))
	if err != nil {
		panic(err)
	}

	sendFile, errorSendFile := api.PostTranslation(stringJson)
	if errorSendFile != nil {
		panic(errorSendFile)
	}
	fmt.Println(sendFile)

	// Get the locales
	locales, err := locales.GetLocales(argsWithProg[0])
	if err != nil {
		panic(err)
	}

	response, err := api.GetTranslationByLocal(argsWithProg[0])
	if err != nil {
		panic(err)
	}
	makJsonData := make(map[string]json.RawMessage)
	errMapJson := json.Unmarshal([]byte(response), &makJsonData)
	if errMapJson != nil {
		panic(errMapJson)
	}

	for _, locale := range locales {
		urlPath := root + string(argsWithProg[3]) + "/" + locale.Code + ".json"
		fmt.Printf("Locale: %s is going to store to %s\n", locale.Code, urlPath)
		WriteInFile(urlPath, string(makJsonData[locale.Code]))
	}

}
