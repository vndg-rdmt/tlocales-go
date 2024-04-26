package tlocales

type ReaderDriver interface {
	Read(key string) ([]RawDictionary, error)
}

type RawDictionary struct {
	Name    string
	Content []byte
}

type Option = func(self *loadContract)

// Holds dictionaries as values by language as keys
type Book map[string]Polyglot

// Holds lexicons by language
type Polyglot map[string]Lexicon

// Holds definitions by key and phrases by value
type Lexicon map[string]string

type Unmarshaller func(b []byte, dest any) error
