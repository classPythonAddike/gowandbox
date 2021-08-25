package gowandbox

var WandBoxUrl string = "https://wandbox.org/api/"

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
