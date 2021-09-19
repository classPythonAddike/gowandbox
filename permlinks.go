package gowandbox

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	Returns information about a program that was run, given the permlink.
	Maps to the `/permlink/link` endpoint.
	If the permlink is not found, and 500 error is returned.
*/
func GetPermLink(link string, ctx context.Context) (GWBPermLink, error) {

	var result GWBPermLink

	client := http.DefaultClient

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,

		WandBoxUrl+"permlink/"+link,
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
		return result, fmt.Errorf("%v Error - %v", resp.StatusCode, string(e))
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	return result, err
}
