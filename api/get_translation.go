package api

import (
	"fmt"
	"io"
	"net/http"
)

const API_URL = "https://localise.biz/api/export/locale/%s.json?key=%s&fallback=en"

func GetTranslationByLocal(locale string) (string, error) {
	if len(locale) == 0 {
		return "", fmt.Errorf("locale is required")
	}
	res, err := http.Get(fmt.Sprintf(API_URL, locale, "hu6pYwo4kq0UL0oeitgr_ugZFeMxcSb3P"))
	if err != nil {
		return "", fmt.Errorf("error on api call to locale error: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error on api call to locale status code: %d", res.StatusCode)
	}
	resData, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		return "", fmt.Errorf("error on read response body: %s", errRead)
	}
	return string(resData), nil

}
