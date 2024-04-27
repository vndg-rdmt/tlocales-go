package tlocales

import "golang.org/x/text/language"

type Locales interface {
	Say(lang string, key string) (string, bool)
}

type localesInstance struct {
	polyglot Polyglot
}

// Say implements Locales.
func (self *localesInstance) Say(lang string, key string) (string, bool) {
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		return "", false
	}

	if tags == nil || len(tags) == 0 {
		return "", false
	}

	var lx Lexicon
	var ok bool
	for i := 0; i < len(tags); i++ {
		if lx, ok = self.polyglot[tags[i].String()]; ok {
			break
		}
	}

	if !ok {
		return "", false
	}

	var msg string
	if msg, ok = lx[key]; ok {
		return msg, true
	}

	return "", false
}
