package gowandbox

import (
	"io"
	"strings"
	"testing"
)

func assertData(message, want, got string, t *testing.T) {
	if want != got {
		t.Errorf("%v - got %v but want %v!", message, want, got)
	}
}

func TestNDExecute(t *testing.T) {
	prog := NewGWBNDProgram()

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
		t.Error(err.Error())
	}

	for {
		msg, err := result.Next()
		if err == io.EOF {
			break
		} else if err == nil {
			switch msg.Type {
			case "StdOut":
				t.Log("Checking stdout")
				assertData("Unexpected output in StdOut", "123\n", msg.Data, t)
			case "ExitCode":
				t.Log("Checking exit code")
				assertData("Unexpected exit code", "0", msg.Data, t)
			}
		} else {
			t.Error(err.Error())
		}
	}

}

func TestNDExecuteTimeout(t *testing.T) {
	prog := NewGWBNDProgram()

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

	_, err := prog.Execute(1)

	if err == nil {
		t.Error("Got no error, but was expecting a timeout!")
	}

	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Error(err.Error())
	}

	t.Log("Request timed out, as expected")
}

func TestNDExecuteBadCompilerError(t *testing.T) {
	prog := NewGWBNDProgram()

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

	_, err := prog.Execute(10000)

	if err == nil {
		t.Error("Got no error, but was expecting a server error!")
	}

	if !strings.Contains(err.Error(), "Internal Server Error") {
		t.Error(err.Error())
	}

	t.Log("Internal Server Error, as expected")
}
