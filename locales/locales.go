package locales

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Locales struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Source  bool   `json:"source"`
	Native  bool   `json:"native"`
	Plurals struct {
		Length   int      `json:"length"`
		Equation string   `json:"equation"`
		Forms    []string `json:"forms"`
	} `json:"plurals"`
	Progress struct {
		Translated   int `json:"translated"`
		Untranslated int `json:"untranslated"`
		Flagged      int `json:"flagged"`
		Words        int `json:"words"`
	} `json:"progress"`
}

const API_URL = "https://localise.biz/api/locales?key=%s"

func GetLocales(token string) ([]Locales, error) {
	if len(token) == 0 {
		return nil, fmt.Errorf("error on api call to locale error: %s", "Token is empty")
	}
	res, err := http.Get(fmt.Sprintf(API_URL, token))
	if err != nil {
		return nil, fmt.Errorf("error on api call to locale error: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error on api call to locale status code: %d", res.StatusCode)
	}
	var locales []Locales
	if err := json.NewDecoder(res.Body).Decode(&locales); err != nil {
		return nil, fmt.Errorf("error on decode response body: %s", err)
	}
	return locales, nil
}
