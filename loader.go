package tlocales

type Loader interface {
	Load(key string, opts ...Option) error
	GetLocales(name string) (Locales, bool)
	GetBook() Book
}

func New() Loader {
	return &instance{
		book: Book{},
	}
}

type instance struct {
	book Book
}

// GetBook implements Locales.
func (self *instance) GetBook() Book {
	return self.book
}

// GetDictionary implements Locales.
func (self *instance) GetLocales(dictname string) (Locales, bool) {
	pl, ok := self.book[dictname]
	if !ok {
		return nil, false
	}

	return &localesInstance{polyglot: pl}, true
}

// Load implements Locales.
func (self *instance) Load(key string, opts ...Option) error {

	lc := loadContract{
		unmarshallers: map[string]Unmarshaller{},
		reader:        nil,
	}

	for i := 0; i < len(opts); i++ {
		opts[i](&lc)
	}

	if lc.reader == nil {
		return &LocalesError{
			Type:    ErrTypeReaderProblem,
			Message: "no reader driver provided",
			Params:  nil,
		}
	}

	if err := lc.loadBook(key, self.book); err != nil {
		return err
	}

	return nil
}
