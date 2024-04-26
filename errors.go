package tlocales

type ErrType string

const (
	ErrTypeReaderProblem     ErrType = "readerProblem"
	ErrTypeMalformedDictName ErrType = "malformedDictName"
	ErrTypeMalformedDict     ErrType = "malformedDict"
)

type LocalesError struct {
	Type    ErrType     `json:"type" yaml:"type"`
	Message string      `json:"err" yaml:"err"`
	Params  interface{} `json:"params,omitempty" yaml:"params"`
}

func (self *LocalesError) Error() string {
	return self.Message
}

type ReaderError struct {
	Driver string `json:"driver" yaml:"driver"`
	Key    string `json:"key" yaml:"key"`
	Error  string `json:"error" yaml:"error"`
}

type MalformedDictNameError struct {
	Driver     string `json:"driver" yaml:"driver"`
	LoadedFrom string `json:"loaded_from" yaml:"loaded_from"`
	Name       string `json:"name" yaml:"name"`
	Error      string `json:"error" yaml:"error"`
}

type MalformedDictError struct {
	Driver     string `json:"driver" yaml:"driver"`
	LoadedFrom string `json:"loaded_from" yaml:"loaded_from"`
	Name       string `json:"name" yaml:"name"`
	Error      string `json:"error" yaml:"error"`
}
