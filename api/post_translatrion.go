package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const URL_API = "https://localise.biz/api/import/json?key=%s&locale=en&ignore-existing=true&tag-absent=obsolete&format=JSON"

func PostTranslation(fileToSend string) (string, error) {
	if len(fileToSend) == 0 {
		return "", fmt.Errorf("filename is required")
	}
	res, err := http.Post(fmt.Sprintf(URL_API, "hu6pYwo4kq0UL0oeitgr_ugZFeMxcSb3P"), "application/json", strings.NewReader(fileToSend))
	if err != nil {
		return "", fmt.Errorf("error on api call to locale error: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error on post api call to locale status code: %d ", res.StatusCode)
	}
	resData, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		return "", fmt.Errorf("error on read response body: %s", errRead)
	}
	fmt.Printf("Update response data %s", resData)
	return fmt.Sprintf("File %s is successfully uploaded", fileToSend), nil
}
