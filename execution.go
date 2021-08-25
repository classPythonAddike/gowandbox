package gowandbox

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

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

	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
