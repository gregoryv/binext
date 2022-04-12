package binext

import "testing"

func TestIsBinary(t *testing.T) {

	cases := []struct {
		v   string
		exp bool
	}{
		{".a", true},
		{".txt", false},
		{".md", false},
	}
	for _, c := range cases {
		t.Run("ext: "+c.v, func(t *testing.T) {
			if IsBinaryString(c.v) != c.exp {
				t.Fail()
			}
			if IsBinary([]byte(c.v)) != c.exp {
				t.Fail()
			}
		})
	}
}
