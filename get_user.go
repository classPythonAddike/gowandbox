package gowandbox

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func GetUser(sessionKey string, timeout int) (GWBUser, error) {
	var result GWBUser

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Get(
		wandBoxUrl + "/user.json?session=" + sessionKey,
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
