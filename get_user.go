package gowandbox

import (
	"encoding/json"
	"errors"
	"fmt"
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
		wandBoxUrl + "template/user.json?session=" + sessionKey,
	)

	if resp.StatusCode != http.StatusOK {
		e, _ := ioutil.ReadAll(resp.Body)
		return result, errors.New(fmt.Sprintf("%v Error - %v", resp.StatusCode, string(e)))
	}

	if err != nil {
		return result, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
