package gowandbox

import (
	"bufio"
)

var defaultUrl string = "https://wandbox.org/api/"

// Url to WandBox API
var WandBoxUrl string = defaultUrl

// Changes the default WandBox URL, useful to use with your own instance of WandBox
func ChangeWandBoxUrl(url string) {
	WandBoxUrl = url
}

// Resets the WandBox URL used
func ResetWandBoxUrl() {
	WandBoxUrl = defaultUrl
}

/*
	Program Struct, used while compiling
	Takes in the filename, and code
*/
type Program struct {
	File string `json:"file"`
	Code string `json:"code"`
}

/*
	GWBProgram struct, to be used when posting to `/compile.json`

	Full compiler list can be seen by using `GetLanguages()`

	Code is a string containing the main code that will be run. It will be
	the entry point to your program.

	Codes may contain other files to include while running/compiling

	Options, CompilerOptionRaw, RuntimeOptionRaw are strings containing comma
	separated values. See github.com/melpon/wandbox/blob/master/kennel2/API.rst#sample-1
	for more ino

	Stdin is a string of newline separated values, which will be passed in as input.

	Passing `true` to SaveCode will generate a permlink that will be returned, and the
	program will be saved to WandBox at that url. Default value is false.
*/
type GWBProgram struct {
	Compiler string `json:"compiler"`

	Code  string    `json:"code"`
	Codes []Program `json:"codes"`

	Options           string `json:"options"`
	CompilerOptionRaw string `json:"compiler-option-raw"`
	RuntimeOptionRaw  string `json:"runtime-option-raw"`

	Stdin string `json:"stdin"`

	SaveCode bool `json:"save"`
}

/*
	GWBNDProgram struct, to be used when posting to `/compile.ndjson`
	The fields are very similar to GWBProgram, but without the SaveCode option
*/
type GWBNDProgram struct {
	Compiler string `json:"compiler"`

	Code  string    `json:"code"`
	Codes []Program `json:"codes"`

	Options           string `json:"options"`
	CompilerOptionRaw string `json:"compiler-option-raw"`
	RuntimeOptionRaw  string `json:"runtime-option-raw"`

	Stdin string `json:"stdin"`
}

/*
	Result of compiling a program (/compile.json endpoint).

	Status and Signal provide information about how the program exited.

	CompilerOutput provides output during compile time, and CompilerError
	provides any errors that occured during compile time.
	CompilerMessage is a combination of both.

	ProgramOutput provides output during runtime, and ProgramError
	provides any errors that occured during run time.
	ProgramMessage is a combination of both.

	Permlink - permlink of the code, only provided if SaveCode was set to true when compiling.
	Url - Url to view the code and output in the browser.
*/
type GWBResult struct {
	Status string `json:"status"`
	Signal string `json:"signal"`

	CompilerOutput  string `json:"compiler_output"`
	CompilerError   string `json:"compiler_error"`
	CompilerMessage string `json:"compiler_message"`

	ProgramOutput  string `json:"program_output"`
	ProgramError   string `json:"program_error"`
	ProgramMessage string `json:"program_message"`

	Permlink string `json:"permlink"`
	Url      string `json:"url"`
}

/*
	Reader to stream data from the `/compile.ndjson` endpoint.
	Data can be obtained by calling the `Next()` method.
*/
type GWBNDReader struct {
	source *bufio.Scanner
}

/*
	Data returned by the `/compile.ndjson` endpoint.
*/
type GWBNDMessage struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

/*
	Language returned by the `/list.json` endpoint.

	CompilerOptionRaw and RuntimeOptionRaw are booleans stating whether they support
	passing of these options to `/compile.json` or `/compile.ndjson`.

	DisplayCompileCommand shows the command that will be run to execute it.

	Name, Version, Language, DisplayName, Templates show information about the language.

	Switches is a list of switches that can be used during compile time.
*/
type GWBLanguage struct {
	CompilerOptionRaw bool `json:"compiler-option-raw"`
	RuntimeOptionRaw  bool `json:"runtime-option-raw"`

	DisplayCompileCommand string `json:"display-compile-command"`

	Switches []struct {
		Default      interface{} `json:"default"`
		Name         string      `json:"name"`
		DisplayFlags string      `json:"display-flags,omitempty"`
		DisplayName  string      `json:"display-name,omitempty"`
	} `json:"switches"`

	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Language    string   `json:"language"`
	DisplayName string   `json:"display-name"`
	Templates   []string `json:"templates"`
}

/*
	Represents a WandBox user.
	Provides the username, and whether the user is logged in or not.
*/
type GWBUser struct {
	Login    bool   `json:"login"`
	Username string `json:"username"`
}

/*
	Provides information about a permlink

	Parameter displays the parameters provided at runtime
	Note that `CreatedAt` represents the time in ISO-8601 format.

	Result shows the output of the compile.
*/
type GWBPermLink struct {
	Parameter struct {
		Compiler string `json:"compiler"`

		Code  string    `json:"code"`
		Codes []Program `json:"codes"`

		Options string `json:"options"`

		Stdin             string `json:"stdin"`
		CompilerOptionRaw string `json:"compiler-option-raw"`
		RuntimeOptionRaw  string `json:"runtime-option-raw"`

		CreatedAt int64 `json:"created_at"`
	} `json:"parameter"`

	Result struct {
		Status string `json:"status"`
		Signal string `json:"signal"`

		CompilerOutput  string `json:"compiler_output"`
		CompilerError   string `json:"compiler_error"`
		CompilerMessage string `json:"compiler_message"`

		ProgramOutput  string `json:"program_output"`
		ProgramError   string `json:"program_error"`
		ProgramMessage string `json:"program_message"`
	} `json:"result"`
}

type GWBTemplate struct {
	Code string `json:"code"`
}
