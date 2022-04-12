package binext

import (
	"fmt"
	"os"
	"testing"
)

func ExampleIsBinary() {
	fmt.Println("somefile.a", IsBinary("somefile.a"))
	fmt.Println("file.o", IsBinary([]byte("file.o")))
	fh, _ := os.Open("ext_test.go")
	fmt.Println(fh.Name(), IsBinary(fh))
	fh.Close()
	fmt.Println("path/to/README.md", IsBinary("path/to/README.md"))
	// output:
	// somefile.a true
	// file.o true
	// ext_test.go false
	// path/to/README.md false
}

func TestIsBinary(t *testing.T) {
	cases := []struct {
		v   interface{}
		exp bool
	}{
		{".a", true},
		{[]byte(".o"), true},
		{os.Stdin, false},

		{".txt", false},
		{".md", false},

		{1, false},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if IsBinary(c.v) != c.exp {
				t.Fail()
			}
		})
	}
}
