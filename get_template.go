package gowandbox

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func GetTemplate(language string, timeout int) (string, error) {

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Get(
		wandBoxUrl + "template/" + language,
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
