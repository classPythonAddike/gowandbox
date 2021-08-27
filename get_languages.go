package gowandbox

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

/*
Returns a list of languages, maps to the `/list.json` endpoint
*/
func GetLanguages(timeout int) ([]GWBLanguage, error) {
	var result []GWBLanguage

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Get(
		WandBoxUrl + "list.json",
	)

	if err != nil {
		return result, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		e, _ := ioutil.ReadAll(resp.Body)
		return result, errors.New(string(e))
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
