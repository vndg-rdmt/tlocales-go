package tlocales

type Dictionary interface {
	Say(lang string, key string) (string, Book)
}

type dict struct {
	lx Lexicon
}

// func (self *dict) Say(lang string, key string) (string, Book) {
// }
