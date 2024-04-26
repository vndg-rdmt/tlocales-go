package driver_fs

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/vndg-rdmt/tlocales-go"
)

func TestDriver(t *testing.T) {
	var err error
	dr := tlocales.New()

	if err = dr.Load("/test/fs",
		tlocales.UseDriver(New("../../")),
		tlocales.RegisterUnmarshaller("json", json.Unmarshal),
	); err != nil {
		t.Error(err)
	}

	var b []byte
	if b, err = json.MarshalIndent(dr.GetBook(), "", "  "); err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
}
