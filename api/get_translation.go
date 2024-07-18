package api

import (
	"fmt"
	"io"
	"net/http"
)

const API_URL = "https://localise.biz/api/export/all.json?fallback=en&no-folding=true&key=%s"

func GetTranslationByLocal(token string) (string, error) {
	res, err := http.Get(fmt.Sprintf(API_URL, token))
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
