package gowandbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetPermLink(link string, timeout int) (GWBPermLink, error) {

	var result GWBPermLink

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Get(
		wandBoxUrl + "permlink/" + link,
	)

	if err != nil {
		return result, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		e, _ := ioutil.ReadAll(resp.Body)
		return result, errors.New(fmt.Sprintf("%v Error - %v", resp.StatusCode, string(e)))
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	return result, err
}
