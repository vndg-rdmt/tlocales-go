package driver_fs

import (
	"fmt"
	"os"
	"path/filepath"

	tlocales_sdk "github.com/vndg-rdmt/tlocales-go/sdk"
)

// Get files from dir entry. Returned filenames are
// contatenated with `prefix` and absolute.
func getFiles(prefix string, dir []os.DirEntry) []string {
	files := make([]string, 0, len(dir))

	for i := 0; i < len(dir); i++ {
		if !dir[i].IsDir() {
			files = append(files, filepath.Join(prefix, dir[i].Name()))
		}
	}

	return files
}

// Reads files provided filenames and stores it's base names
// with its contents in `tlocales_sdk.RawDictionary` struct.
func readFiles(files []string) ([]tlocales_sdk.RawDictionaryInterface, error) {
	res := make([]tlocales_sdk.RawDictionaryInterface, len(files))

	var buffer []byte
	var err error

	for i := 0; i < len(files); i++ {
		if buffer, err = os.ReadFile(files[i]); err != nil {
			return nil, fmt.Errorf("cannot read file %s, %w", files[i], err)
		}

		res[i] = tlocales_sdk.RawDictionaryInterface{
			Name:    filepath.Base(files[i]),
			Content: buffer,
		}
	}

	return res, nil
}
