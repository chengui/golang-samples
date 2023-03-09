package echo

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		newl bool
		args []string
		want string
	}{
		{
			true,
			[]string{"hello", "world"},
			"hello world",
		},
		{
			false,
			[]string{"hello", "world"},
			"hello world\n",
		},
	}
	for _, test := range tests {
		var buf bytes.Buffer
		newlineFlag = test.newl
		if err := echo(&buf, test.args); err != nil {
			t.Errorf("error: %v\n", err)
		}
		if got := buf.String(); got != test.want {
			t.Errorf("%s != %s", got, test.want)
		}
	}
}
