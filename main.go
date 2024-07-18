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
	var token string
	var pathToGetJson string
	var locale string
	var pathToStoreJson string

	if len(argsWithProg) != 4 {
		fmt.Println("Please provide the token")
		fmt.Scan(&token)
		fmt.Println("Please provide the Path to get the json")
		fmt.Scan(&pathToGetJson)
		fmt.Println("Please provide the locale")
		fmt.Scan(&locale)
		fmt.Println("Please provide the path to store the json")
		fmt.Scan(&pathToStoreJson)
	} else {
		fmt.Println("-------------------")
		fmt.Printf("Arguments Token: %s\n", argsWithProg[0])
		fmt.Println("-------------------")
		fmt.Printf("Arguments Path to get the json: %s\n", argsWithProg[1])
		fmt.Println("-------------------")
		fmt.Printf("Arguments locale: %s\n", argsWithProg[2])
		fmt.Println("-------------------")
		fmt.Printf("Arguments path to store the json: %s \n", argsWithProg[3])
		fmt.Println("-------------------")
		token = argsWithProg[0]
		pathToGetJson = argsWithProg[1]
		locale = argsWithProg[2]
		pathToStoreJson = argsWithProg[3]
	}
	if locale == "" || token == "" || pathToGetJson == "" || pathToStoreJson == "" {
		panic("Please provide the token, path to get the json, locale and path to store the json")
	}

	stringJson, err := filemanage.ReadTextFile(root + string(pathToGetJson))
	if err != nil {
		panic(err)
	}

	sendFile, errorSendFile := api.PostTranslation(stringJson, locale)
	if errorSendFile != nil {
		panic(errorSendFile)
	}
	fmt.Println(sendFile)

	// Get the locales
	locales, err := locales.GetLocales(token)
	if err != nil {
		panic(err)
	}

	response, err := api.GetTranslationByLocal(token)
	if err != nil {
		panic(err)
	}
	makJsonData := make(map[string]json.RawMessage)
	errMapJson := json.Unmarshal([]byte(response), &makJsonData)
	if errMapJson != nil {
		panic(errMapJson)
	}

	for _, locale := range locales {
		urlPath := root + string(pathToStoreJson) + "/" + locale.Code + ".json"
		fmt.Printf("Locale: %s is going to store to %s\n", locale.Code, urlPath)
		WriteInFile(urlPath, string(makJsonData[locale.Code]))
	}

}
