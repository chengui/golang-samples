package wc

import (
	"bytes"
	"strings"
	"testing"
)

func TestWc(t *testing.T) {
	tests := []struct {
		option Option
		text   string
		path   string
		want   string
	}{
		{
			Option{false, true, false, false},
			"hello world\n\nhello world 2\n",
			"hello.txt",
			"       3 hello.txt\n",
		},
		{
			Option{false, false, true, false},
			"hello world\n\nhello world 2\n",
			"hello.txt",
			"       5 hello.txt\n",
		},
		{
			Option{false, false, false, true},
			"hello world\n\nhello world 2\n",
			"hello.txt",
			"      27 hello.txt\n",
		},
		{
			Option{false, false, false, false},
			"hello world\n\nhello world 2\n",
			"hello.txt",
			"       3       5      27 hello.txt\n",
		},
		{
			Option{false, true, true, true},
			"hello world\n\nhello world 2\n",
			"hello.txt",
			"       3       5      27 hello.txt\n",
		},
	}
	for _, test := range tests {
		var buf bytes.Buffer
		opt = test.option
		f := strings.NewReader(test.text)
		if _, err := wc(&buf, f, test.path); err != nil {
			t.Errorf("error: %v", err)
		}
		if got := buf.String(); got != test.want {
			t.Errorf("%q != %q", got, test.want)
		}
	}
}
