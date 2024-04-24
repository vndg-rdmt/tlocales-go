package tlocales

import (
	"fmt"
	"strings"
)

// Reads dictionary targe name into the `lcf`
func readName(name string, lcf *LocalesObject) error {
	localsConfig := strings.SplitN(name, ".", 3)
	if len(localsConfig) != 3 {
		return fmt.Errorf("expected to see name 'name.lang.extension'")
	}

	var attrname string
	var attrp *string

	for i := 0; i < len(localsConfig); i++ {
		switch i {
		case 0:
			attrname = "name"
			attrp = &lcf.Dictionary
		case 1:
			attrname = "language"
			attrp = &lcf.Lang
		case 2:
			attrname = "extension"
			attrp = &lcf.Extension
		}

		if localsConfig[i] == "" {
			return fmt.Errorf("attribute %s at name position %d cannot be empty", attrname, i)
		}
		*attrp = localsConfig[i]
	}
	return nil
}
