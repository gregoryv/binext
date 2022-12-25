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
		{"file.a", true},
		{[]byte("path/to/file.o"), true},
		{`.\..\windows.dll`, true},
		{`image.png`, true},
		{`image.PNG`, true},

		{os.Stdin, false},
		{"plain.txt", false},
		{"./../readme.md", false},

		{1, false}, // bad type
		{nil, false},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.v), func(t *testing.T) {
			got := IsBinary(c.v)
			if got != c.exp {
				t.Error("got", got)
			}
		})
	}
}
