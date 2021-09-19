package gowandbox

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func NewGWBProgram() *GWBProgram {
	// Returns new GWBProgram struct
	return &GWBProgram{}
}

/*
Method to execute a GWBProgram

If no errors ocurred, the result is returned in the form of a GWBResult struct.
If the response code is not 200, an error is returned.

Maps to the `/compile.json` endpoint
*/
func (g *GWBProgram) Execute(ctx context.Context) (GWBResult, error) {

	data, err := json.Marshal(g)
	var result GWBResult

	if err != nil {
		return result, err
	}

	client := http.DefaultClient

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		WandBoxUrl+"compile.json",
		bytes.NewBuffer(data),
	)

	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "application/json")

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
