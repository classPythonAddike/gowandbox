# Compiling

GoWandBox (which wraps around the WandBox API) provides two methods to compile your program -

## With GWBProgram

You can compile code by using the `GWBProgram.Execute()` method.

Make sure you've imported `github.com/classPythonAddike/gowandbox` as `gwb`

First initialise a new GWBProgram - 
```go
prog := gwb.NewGWBProgram()

prog.Code = `
import gwbutil
gwbutil.say()
gwbutil.say()
` // Main program

// Additional files to use
prog.Codes = []gwb.Program{
	{
		"gwbutil.py",
		"def say(): print(input())",
	},
}

// Show warnings
prog.Options = "warning"
prog.Compiler = "cpython-3.8.0" // Use cpython 3.8
prog.Stdin = "123\n456" // Input for the program
prog.SaveCode = true // Save it to a permlink
```

Execute it with `.Execute()`
```go
result, err := prog.Execute(
		context.WithTimeout(context.Background()),
		10000 * time.Millisecond
	) // 10000 milliseconds is the timeout

if err != nil {
	log.Fatal(err)
	/*
		`err` is not nil if the response status code is not 200,
		if the request timed out,
		or an intnernal error ocurred in the wrapper
	*/
}
```

You can access information like errors, exit code and output from the `GWBResult` struct.
```go
fmt.Printf("Output during compile time - \"%v\"\n", result.CompilerOutput)
fmt.Printf("Errors during compile time - \"%v\"\n", result.CompilerError)

fmt.Printf("Output of program - \"%v\"\n", result.ProgramOutput)
fmt.Printf("Errors during runtime - \"%v\"\n", result.ProgramError)

fmt.Printf("Exit code - \"%v\"\n", result.Status)
fmt.Printf("Signal - \"%v\"\n", result.Signal)

fmt.Printf("View the program in your browser at \"%v\"\n", result.Url)
fmt.Printf("Permlink to this program - \"%v\"\n", result.Permlink)
```

```
$ go run main.go
Output during compile time - ""
Errors during compile time - ""
Output of program - "123
456
"
Errors during runtime - ""
Exit code - "0"
Signal - ""
View the program in your browser at "https://wandbox.org/permlink/kOJgNvv8flGlApCA"
Permlink to this program - "kOJgNvv8flGlApCA"
```

## With GWBNDProgram

WandBox offers an alternate method to compile your code - by using the `/compile.ndjson` endpoint. While this method is ideal for "streaming" the response, as it returns `NDJSON`, it is not as verbose as the former method.

GoWandBox's API for GWBProgram and GWBNDProgram are very similar.

First, declare your `GWBNDProgram` -
```go
prog := gwb.NewGWBNDProgram()

prog.Code = `
import gwbutil, time
gwbutil.say()
time.sleep(5)
gwbutil.say()
`
prog.Codes = []gwb.Program{
	{
		"gwbutil.py",
		"def say(): print(input())",
	},
}
prog.Options = "warning"
prog.Compiler = "cpython-3.8.0"
prog.Stdin = "123\n456"
```

Then, execute your code -
```go
result, err := prog.Execute(
		context.WithTimeout(context.Background()),
		10000 * time.Millisecond
	) // 10000 milliseconds is the timeout

if err != nil {
	log.Fatal(err)
}
```

You can obtain the data with the `.Next()` method of the `GWBNDReader` -
```go
for {
	msg, err := result.Next() // msg is of the type GWBNDMessage

	if err == io.EOF {
		break
	} else if err == nil {
		fmt.Printf("Type - %v, Data - %v", msg.Type, msg.Data)
	} else {
		log.Fatal(err)
	}
}
```

```
$ go run main.go
Type - Control, Data - Start
Type - StdOut, Data - 123
Type - StdOut, Data - 456
Type - ExitCode, Data - 0
Type - Control, Data - Finish
```
