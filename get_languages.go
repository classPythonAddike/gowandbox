package gowandbox

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func GetLanguages(timeout int) ([]GWBLanguage, error) {
	var result []GWBLanguage

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Get(
		wandBoxUrl + "list.json",
	)

	if err != nil {
		return result, err
	}

	if resp.StatusCode != http.StatusOK {
		e, _ := ioutil.ReadAll(resp.Body)
		return result, errors.New(string(e))
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
