package gowandbox

var wandBoxUrl string = "https://wandbox.org/api/"

type Program struct {
	File string `json:"file"`
	Code string `json:"code"`
}

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

type GWBUser struct {
	Login    bool   `json:"login"`
	Username string `json:"username"`
}
