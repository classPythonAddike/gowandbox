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

**Note**: This wrapper supports most of the WandBox API's endpoints, except for the `POST /permlink`. This was because, I couldn't figure out how to work with the endpoint - I kept winding up with `500 - Internal Server Errors`. I'm not sure whether it was a bug in the test data that I was using, or if its a bug in the API itself. Either way, that endpoint has not been implemented in this wrapper, and I apologise for any inconveniences. Feel free to make a PR, or open an issue if you're able to figure it out!

## Installation

```
$ go get github.com/classPythonAddike/gowandbox
```


## Quick Start

```go

package main

import (
	"log"
	"context"
	"time"

	gwb "github.com/classPythonAddike/gowandbox"
)

func main() {

	prog := gwb.NewGWBProgram()

	prog.Code = "print('Hello, World!')" // Code for the main program
	prog.Compiler = "cpython-3.8.0"      // Use cpython, 3.8 to run your code

	ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	result, err := prog.Execute(ctx) // 10 seconds is the timeout

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Output: %v", result.ProgramOutput)
}

```

```
$ go run main.go
Output: Hello, World!

```
