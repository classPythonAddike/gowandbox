package gowandbox

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"
)

func assert(expected, got string, t *testing.T) {
	if expected != got {
		t.Errorf("Expected - %v, but got %v!", expected, got)
	}
}

func TestExecution(t *testing.T) {

	prog := NewGWBProgram()

	prog.Code = "import gwbutil\n\ngwbutil.say()\ngwbutil.say()"
	prog.Codes = []Program{
		{
			"gwbutil.py",
			"def say(): print(input())",
		},
	}
	prog.Options = "warning"
	prog.Compiler = "cpython-3.8.0"
	prog.Stdin = "123\n456"

	result, err := prog.Execute(context.Background())

	if err != nil {
		t.Error(err.Error())
	}

	assert("123\n456\n", result.ProgramOutput, t)

	log.Printf("Got output - %v", result.ProgramOutput)
	log.Println("Comparing errors, compiler output")

	assert("", result.CompilerError, t)
	assert("", result.ProgramError, t)
	assert("", result.CompilerOutput, t)
}

func TestExecutionTimeout(t *testing.T) {

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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	_, err := prog.Execute(ctx)

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Error(err.Error())
	}

	t.Log("Request timed out, as expected")
}

func TestExecutionBadCompiler(t *testing.T) {

	prog := NewGWBProgram()

	prog.Code = "import gwbutil\n\ngwbutil.say()"
	prog.Codes = []Program{
		{
			"gwbutil.py",
			"def say(): print(input())",
		},
	}
	prog.Options = "warning"
	prog.Compiler = "abc"
	prog.Stdin = "123"

	_, err := prog.Execute(context.Background())

	if err == nil {
		t.Error("Got no error, but was expecting one!")
	}

	if !strings.Contains(err.Error(), "Internal Server Error") {
		t.Error(err.Error())
	}

	t.Log("Program was not compiled, as expected")
}
