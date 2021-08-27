package gowandbox

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Returns new GWBProgram struct, after filling it out with defaults
func NewGWBProgram() GWBProgram {
	return GWBProgram{
		Compiler:          "",
		Code:              "",
		Codes:             []Program{},
		Options:           "",
		CompilerOptionRaw: "",
		RuntimeOptionRaw:  "",
		Stdin:             "",
		SaveCode:          false,
	}
}

/*
Method to execute a GWBProgram

If no errors ocurred, the result is returned in the form of a GWBResult struct.
If the response code is not 200, an error is returned.

Maps to the `/compile.json` endpoint
*/
func (g *GWBProgram) Execute(timeout int) (GWBResult, error) {

	data, err := json.Marshal(g)
	var result GWBResult

	if err != nil {
		return result, err
	}

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Post(
		WandBoxUrl+"compile.json",
		"application/json",
		bytes.NewBuffer(data),
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
