package gowandbox

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Returns new GWBProgram struct, after filling it out with defaults.
func NewGWBNDProgram() GWBNDProgram {
	return GWBNDProgram{
		Compiler:          "",
		Code:              "",
		Codes:             []Program{},
		Options:           "",
		CompilerOptionRaw: "",
		RuntimeOptionRaw:  "",
		Stdin:             "",
	}
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
func (g *GWBNDProgram) Execute(timeout int) (GWBNDReader, error) {

	data, err := json.Marshal(g)
	var result GWBNDReader

	if err != nil {
		return result, err
	}

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := client.Post(
		WandBoxUrl+"compile.ndjson",
		"application/json",
		bytes.NewBuffer(data),
	)

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
