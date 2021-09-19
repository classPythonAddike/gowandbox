package gowandbox

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

/*
	Returns the template for a given language.
	If the language is not found, an error is returned (bad template).
	Maps to the `/template/:template` endpoint
*/
func GetTemplate(language string, ctx context.Context) (GWBTemplate, error) {

	client := http.DefaultClient

	templ := GWBTemplate{}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		WandBoxUrl+"template/"+language,
		nil,
	)

	if err != nil {
		return templ, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return templ, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		e, _ := ioutil.ReadAll(resp.Body)
		return templ, errors.New(string(e))
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&templ)
	return templ, err
}
