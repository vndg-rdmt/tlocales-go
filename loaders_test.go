package tlocales

import (
	"testing"
)

type caseReadname struct {
	Name      string
	Malformed bool
}

func getCasesReadName() []caseReadname {
	return []caseReadname{
		{
			Name:      "dwqdqwdqwdqwdwq",
			Malformed: true,
		},
		{
			Name: "dwq.r.qqq",
		},
		{
			Name: "name.lang.ext",
		},
		{
			Name: "dwuqbdyuwqbvdyqwb.en.json",
		},
		{
			Name:      "dwuqbdyuwqbvdyqwb..json",
			Malformed: true,
		},
		{
			Name:      ".en.json",
			Malformed: true,
		},
		{
			Name:      "..",
			Malformed: true,
		},
		{
			Name:      ".dwq.",
			Malformed: true,
		},
		{
			Name:      "dwqdqw..",
			Malformed: true,
		},
	}
}

func TestReadName(t *testing.T) {
	var lcfMock LocalesObject
	var err error
	for i, tcs := range getCasesReadName() {
		err = readName(tcs.Name, &lcfMock)
		if (err != nil) != tcs.Malformed {
			t.Errorf("case [%d], malformed ? - %t, error - %v", i, tcs.Malformed, err)
		}
	}
}
