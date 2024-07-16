package filemanage

import (
	"fmt"
	"os"
)

func WriteInFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0777)
	if err != nil {
		return err
	}
	fmt.Println("File successfully is write")
	return nil
}
