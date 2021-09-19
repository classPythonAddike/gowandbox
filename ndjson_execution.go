package gowandbox

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// Returns new GWBProgram struct
func NewGWBNDProgram() *GWBNDProgram {
	return &GWBNDProgram{}
}

/*
	Method to read ndjson.
	Returns a GWNDBMessage struct, which provides the type of the message, and data.
	On reaching the end, an `io.EOF` error is returned.
*/
func (r *GWBNDReader) Next() (GWBNDMessage, error) {
	if !r.source.Scan() {
		return GWBNDMessage{}, io.EOF
	}

	data := []byte(r.source.Text())

	if err := r.source.Err(); err != nil {
		return GWBNDMessage{}, err
	}

	result := GWBNDMessage{}
	err := json.Unmarshal(data, &result)

	return result, err
}

/*
Method to execute a GWBProgram

If no errors ocurred, the result is returned in the form of a GWBNDReader struct.
If the response code is not 200, an error is returned.

Maps to the `/compile.ndjson` endpoint
*/
func (g *GWBNDProgram) Execute(ctx context.Context) (GWBNDReader, error) {

	data, err := json.Marshal(g)
	var result GWBNDReader

	if err != nil {
		return result, err
	}

	client := http.DefaultClient

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,

		WandBoxUrl+"compile.ndjson",

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
		e, _ := ioutil.ReadAll(resp.Body)
		return result, errors.New(string(e))
	}

	result.source = bufio.NewScanner(resp.Body)
	return result, err
}
