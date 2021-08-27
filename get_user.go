package gowandbox

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	Returns a user, given the session key (provided by WandBox).
	If the session key is invalid, a blank string is returned for the username.
	Maps to the `/user.json` endpoint.
*/
func GetUser(sessionKey string, timeout int) (GWBUser, error) {
	var result GWBUser

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Get(
		WandBoxUrl + "/user.json?session=" + sessionKey,
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
