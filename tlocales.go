package tlocales

type LocalesObject struct {
	Dictionary string
	Lang       string
	Extension  string
}

// Holds dictionaries as values by their names as keys
type Book map[string]Dictionary

// Holds definitions by key and phrases by value
type Dictionary map[string]string
