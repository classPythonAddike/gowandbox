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
