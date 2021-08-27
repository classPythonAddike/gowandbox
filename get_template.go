package gowandbox

import (
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
func GetTemplate(language string, timeout int) (string, error) {

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Get(
		WandBoxUrl + "template/" + language,
	)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		e, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(string(e))
	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	return string(res), err
}
