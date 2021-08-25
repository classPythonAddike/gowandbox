package gowandbox

import (
	"testing"
)

func TestMain(t *testing.T) {

	p := GWBProgram{
		Code:     "print(input())",
		Options:  "warning",
		Compiler: "cpython-3.8.0",
		Stdin:    "123",

		CompilerOptionRaw: "",
		RuntimeOptionRaw:  "",
		Codes:             []Program{},
		SaveCode:          false,
	}

	result, err := p.Execute(10000)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
