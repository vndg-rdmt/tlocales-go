package tlocales

type Locales interface {
	Load(key string, opts ...Option) error
	GetDictionary(name string) (Dictionary, bool)
	GetBook() Book
}

func New() Locales {
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
func (self *instance) GetDictionary(name string) (Dictionary, bool) {
	panic("")
	// lx, ok := self.book[name]
	// if !ok {
	// 	return nil, false
	// }

	// return dict{lx: lx}, true
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
