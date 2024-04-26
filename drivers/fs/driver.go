package driver_fs

import (
	"fmt"
	"os"
	"path/filepath"

	tlocales_sdk "github.com/vndg-rdmt/tlocales-go/sdk"
)

func New(root string) tlocales_sdk.DriverInterface {
	return &fs_driver{
		root: root,
	}
}

type fs_driver struct {
	// filapath prefix which contacentated with
	// key to read dictionary.
	root string
}

// Read implements tlocales_sdk.ReaderDriver.
func (self *fs_driver) Read(key string) ([]tlocales_sdk.RawDictionaryInterface, error) {
	dest := filepath.Join(self.root, key)

	var de []os.DirEntry
	var err error

	if de, err = os.ReadDir(dest); err != nil {
		return nil, fmt.Errorf("cannot read directory, %w", err)
	}

	return readFiles(getFiles(dest, de))
}
