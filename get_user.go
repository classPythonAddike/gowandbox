package gowandbox

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

/*
	Returns a user, given the session key (provided by WandBox).
	If the session key is invalid, a blank string is returned for the username.
	Maps to the `/user.json` endpoint.
*/
func GetUser(sessionKey string, ctx context.Context) (GWBUser, error) {
	var result GWBUser

	client := http.DefaultClient

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		WandBoxUrl+"/user.json?session="+sessionKey,
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
