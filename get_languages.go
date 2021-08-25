package gowandbox

import (
	"encoding/json"
	"errors"
	"fmt"
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

	if resp.StatusCode != http.StatusOK {
		e, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(fmt.Sprintf("%v Error - %v", resp.StatusCode, string(e)))
	}

	if err != nil {
		return result, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
