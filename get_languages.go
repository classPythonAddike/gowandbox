package gowandbox

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

/*
Returns a list of languages, maps to the `/list.json` endpoint
*/
func GetLanguages(ctx context.Context) ([]GWBLanguage, error) {
	var result []GWBLanguage

	client := http.DefaultClient

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		WandBoxUrl+"list.json",
		nil,
	)

	if err != nil {
		return result, err
	}

	resp, err := client.Do(req)

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
