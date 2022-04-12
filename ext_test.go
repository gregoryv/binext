package binext

import (
	"os"
	"testing"
)

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
