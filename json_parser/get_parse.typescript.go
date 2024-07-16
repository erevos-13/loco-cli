package jsonparser

import (
	"fmt"
	"strings"

	"voutsaridis.com/loco-translate-cli/filemanage"
)

func GetParsePathToJson(pathToSend string) (string, error) {
	response, err := filemanage.ReadTextFile(pathToSend)
	if err != nil {
		return "", err
	}

	pattern := "export const TRANSLATE = "
	index := strings.Index(response, pattern)
	if index != -1 {
		content := response[index+len(pattern):]

		// Find the start of the JSON object
		jsonStartIndex := strings.Index(content, "{")
		if jsonStartIndex == -1 {
			return "", fmt.Errorf("valid JSON object not found")
		}
		content = content[jsonStartIndex:]

		return content, nil
	} else {
		return "", fmt.Errorf("pattern not found")
	}
}
