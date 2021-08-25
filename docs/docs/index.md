# Welcome to MkDocs

For full documentation visit [mkdocs.org](https://www.mkdocs.org).

## Installation

* `go get github.com/classPythonAddike/gowandbox`

## Basic Usage

```golang

package main

import (
	"fmt"
	gwb "github.com/classPythonAddike/gowandbox"
)

function main() {
	
	prog := GWBProgram{
		Code:     "print(input())",
		Options:  "warning",
		Compiler: "cpython-3.8.0",
		Stdin:    "123",

		CompilerOptionRaw: "",
		RuntimeOptionRaw:  "",
		Codes:             []Program{},
		SaveCode:          false,
	}

	result, err := prog.Execute(10000) // Timeout is 10,000 milliseconds

	if err != nil {
		fmt.Fatal(err)
	}

	fmt.Sprintf("Output: %v\n", result.CompilerOutput)
}

```
