package gowandbox

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	Returns the template for a given language.
	If the language is not found, an error is returned (bad template).
	Maps to the `/template/:template` endpoint
*/
func GetTemplate(language string, timeout int) (GWBTemplate, error) {

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	templ := GWBTemplate{}

	resp, err := client.Get(
		WandBoxUrl + "template/" + language,
	)

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
