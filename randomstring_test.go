package randomstring

import (
	"testing"
	"unicode"
)

func TestBasic(t *testing.T) {
	str := Generate(12)
	if str == "" || len(str) != 12 {
		t.FailNow()
	}
}

func TestNumCharset(t *testing.T) {
	rsg, err := NewRandomStringGenerator(CharsetNum)
	if err != nil {
		t.Error(err)
	}
	str := rsg.Generate(6)
	if str == "" || len(str) != 6 {
		t.FailNow()
	}
	for _, b := range str {
		if !unicode.IsDigit(b) {
			t.FailNow()
		}
	}
}
