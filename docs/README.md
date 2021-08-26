<h1 align="center">GoWandBox</h1>

<div align="center">
  <img src="https://img.shields.io/github/languages/top/classPythonAddike/gowandbox">
  <img src="https://pkg.go.dev/badge/github.com/classPythonAddike/gowandbox.svg" alt="Go Reference">
  <img src="https://goreportcard.com/badge/github.com/classPythonAddike/gowandbox">
  <img src="https://sourcegraph.com/github.com/classPythonAddike/gowandbox/-/badge.svg">
  <img src="https://www.codetriage.com/classpythonaddike/gowandbox/badges/users.svg">
  <img src="https://img.shields.io/github/license/classPythonAddike/gowandbox">
</div>

<br>

A simple wrapper for the WandBox API, written in Golang!

## Installation

`$ go get github.com/classPythonAddike/gowandbox`


## Usage

```go
package main

import (
    "fmt"
    gwb "github.com/classPythonAddike/gowandbox"
)

func main() {

	prog := NewGWBProgram()

	prog.Code = "import gwbutil\n\ngwbutil.say()\ngwbutil.say()" // Code for the main program

	// Other files to create
	prog.Codes = []Program{
		{
			"gwbutil.py",
			"def say():\nprint(input())",
		},
	}

	prog.Options = "warning" // Show warnings
	prog.Compiler = "cpython-3.8.0" // Use cpython, 3.8 to run your code
	prog.Stdin = "123\n456" // Specify input for the program

	result, err := prog.Execute(10000) // 10000 milliseconds is the timeout

    if err != nil {
        fmt.Fatal(err)
    }

    fmt.Sprintf("Output: %v\n", result.ProgramOutput)
	/*
		Use result.CompilerOutput for output during compile time
		Use result.CompilerError for errors during compile time
		Use result.CompilerMessage for both put together

		Use result.ProgramOutput for output during runtime
		Use result.ProgramError for errors during runtime
		Use result.ProgramMessage for both put together
	*/
}
```
