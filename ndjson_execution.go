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
		wandBoxUrl+"compile.ndjson",
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
