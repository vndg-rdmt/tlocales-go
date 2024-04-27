package tlocales

import (
	"fmt"
	"reflect"
	"strings"
)

func (self Book) define(lcf *localesBuffer, dict Lexicon) {
	var pl Polyglot
	var ok bool

	if pl, ok = self[lcf.dict]; !ok {
		pl = make(Polyglot)
		self[lcf.dict] = pl
	}

	pl[lcf.lang] = dict
}

type loadContract struct {
	unmarshallers map[string]Unmarshaller
	reader        ReaderDriver
}

type localesBuffer struct {
	// name of a dictionary.
	dict string
	// language code of a dictionary.
	lang string
	// extension of a dictionary object, which
	// determines what marshaller shoud be used
	// to parse its content.
	ext string
}

// Shoud a name of a reigistered driver no matter
// of realisation.
//
// Currently reflection is used, by shoud be changed
// to allow users to specify driver name the themself.
func (self *loadContract) driverName() string {
	return reflect.Indirect(reflect.ValueOf(self.reader)).Type().Name()
}

// Loads dictionaries by key for current
// reader driver to a Book `bk`.
func (self *loadContract) loadBook(key string, bk Book) *LocalesError {
	var lu []RawDictionary
	var err error

	if lu, err = self.reader.Read(key); err != nil {
		return &LocalesError{
			Type:    ErrTypeReaderProblem,
			Message: "cannot read dictionaries",
			Params: ReaderError{
				Driver: self.driverName(),
				Key:    key,
				Error:  err.Error(),
			},
		}
	}

	for i := 0; i < len(lu); i++ {
		var lcf localesBuffer
		if err = parseName(lu[i].Name, &lcf); err != nil {
			return &LocalesError{
				Type:    ErrTypeMalformedDictName,
				Message: "cannot parse dictionary name",
				Params: MalformedDictNameError{
					Driver:     self.driverName(),
					LoadedFrom: key,
					Name:       lu[i].Name,
					Error:      err.Error(),
				},
			}
		}

		var dict Lexicon
		if dict, err = parseDict(&lcf, lu[i].Content, self.unmarshallers); err != nil {
			return &LocalesError{
				Type:    ErrTypeMalformedDict,
				Message: "cannot parse dictionary",
				Params: MalformedDictError{
					Driver:     self.driverName(),
					LoadedFrom: key,
					Name:       lu[i].Name,
					Error:      err.Error(),
				},
			}
		}

		bk.define(&lcf, dict)
	}

	return nil
}

// Parses dictionary name into `lcf`
func parseName(name string, lcf *localesBuffer) error {
	localesConfig := strings.SplitN(name, ".", 3)
	if len(localesConfig) != 3 {
		return fmt.Errorf("name must follow pattern 'name.lang.extension'")
	}

	var attrname string
	var attrp *string

	for i := 0; i < len(localesConfig); i++ {
		switch i {
		case 0:
			attrname = "name"
			attrp = &lcf.dict
		case 1:
			attrname = "language"
			attrp = &lcf.lang
		case 2:
			attrname = "extension"
			attrp = &lcf.ext
		}

		if localesConfig[i] == "" {
			return fmt.Errorf("name attribute %s at position %d cannot be empty", attrname, i)
		}
		*attrp = localesConfig[i]
	}
	return nil
}

// Parses dictionary dictionary content
func parseDict(lcf *localesBuffer, content []byte, umumrs map[string]Unmarshaller) (Lexicon, error) {
	var unmarshal Unmarshaller
	var ok bool

	if unmarshal, ok = umumrs[lcf.ext]; !ok {
		return nil, fmt.Errorf("unsupported extension '%s', marshaller for such type not registered", lcf.ext)
	}

	var dict Lexicon = make(Lexicon)
	if err := unmarshal(content, &dict); err != nil {
		return nil, fmt.Errorf("cannot unmarshal content, error - %v", err)
	}

	return dict, nil
}
