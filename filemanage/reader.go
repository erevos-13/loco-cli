package filemanage

import "os"

func ReadTextFile(readFile string) (string, error) {
	content, err := os.ReadFile(readFile)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
