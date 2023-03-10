package dirname

import (
	"bytes"
	"testing"
)

func TestDirname(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{
			[]string{"/a/b/cde"},
			"/a/b\n",
		},
		{
			[]string{"/a/b/cde", "de"},
			"/a/b\n.\n",
		},
		{
			[]string{"/a/b/cde", "de", "../de"},
			"/a/b\n.\n..\n",
		},
	}
	for i, test := range tests {
		var buf bytes.Buffer
		if err := dirname(&buf, test.args); err != nil {
			t.Errorf("error: %v", err)
		}
		if got := buf.String(); got != test.want {
			t.Errorf("#%d %q != %q", i, got, test.want)
		}
	}
}
