package tlocales_sdk

import "github.com/vndg-rdmt/tlocales-go"

// Implement this interface for loading dictionaries
// so it can be fs, object-store and etc.
type DriverInterface = tlocales.ReaderDriver

type RawDictionaryInterface = tlocales.RawDictionary
