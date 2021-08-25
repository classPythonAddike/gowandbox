package gowandbox

import (
	"errors"
	"fmt"
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

	if resp.StatusCode != http.StatusOK {
		e, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(fmt.Sprintf("%v Error - %v", resp.StatusCode, string(e)))
	}

	if err != nil {
		return "", err
	}

	res, err := ioutil.ReadAll(resp.Body)
	return string(res), err
}
