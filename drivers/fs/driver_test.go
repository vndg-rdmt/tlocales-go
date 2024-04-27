package driver_fs

import (
	"encoding/json"
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

	lc, ok := dr.GetLocales("something")
	if !ok {
		t.Error("not found locales")
	}

	if msg, ok := lc.Say("ch", "key1"); !ok {
		t.Error("msg not found")
	} else {
		t.Log(msg)
	}
}
