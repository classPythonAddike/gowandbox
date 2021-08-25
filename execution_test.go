package gowandbox

import (
	"log"
	"testing"
)

func Assert(expected, got string, t *testing.T) {
	if expected != got {
		t.Errorf("Expected - %v, but got %v!", expected, got)
	}
}

func TestMain(t *testing.T) {

	prog := NewGWBProgram()

	prog.Code = "import gwbutil\n\ngwbutil.say()"
	prog.Codes = []Program{
		{
			"gwbutil.py",
			"def say(): print(input())",
		},
	}
	prog.Options = "warning"
	prog.Compiler = "cpython-3.8.0"
	prog.Stdin = "123"

	result, err := prog.Execute(10000)

	if err != nil {
		t.Error(err)
	}

	Assert("123\n", result.ProgramOutput, t)

	log.Printf("Got output - %v", result.ProgramOutput)
	log.Println("Comparing errors, compiler output")

	Assert("", result.CompilerError, t)
	Assert("", result.ProgramError, t)
	Assert("", result.CompilerOutput, t)
}
