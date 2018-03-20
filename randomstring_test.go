package randomstring

import (
	"fmt"
	"log"
	"testing"
	"unicode"

	"github.com/nouney/randomstring"
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

// Use the default generator.
func ExampleRandomStringGenerator() {
	str := randomstring.Generate(16)
	fmt.Println("Random string:", str)
}

// Create a generator with a specific charset.
func ExampleRandomStringGenerator_charset() {
	rsg, err := randomstring.NewGenerator(randomstring.CharsetNum)
	if err != nil {
		log.Fatal(err)
	}
	str := rsg.Generate(16)
	fmt.Println("Random string:", str)
}
