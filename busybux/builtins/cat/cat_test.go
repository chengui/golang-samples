package cat

import (
	"bytes"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	tests := []struct {
		nflag bool
		text  string
		want  string
	}{
		{
			true,
			"hello world\n\nhello world 2\n",
			"1 hello world\n2 \n3 hello world 2\n",
		},
		{
			false,
			"hello world\n\nhello world 2\n",
			"hello world\n\nhello world 2\n",
		},
		{
			true,
			"hello world\n\nhello world 2",
			"1 hello world\n2 \n3 hello world 2",
		},
		{
			false,
			"hello world\n\nhello world 2",
			"hello world\n\nhello world 2",
		},
	}
	for i, test := range tests {
		var buf bytes.Buffer
		numberFlag = test.nflag
		f := strings.NewReader(test.text)
		if err := cat(&buf, f); err != nil {
			t.Errorf("error: %v", err)
		}
		if got := buf.String(); got != test.want {
			t.Errorf("#%d %q != %q", i, got, test.want)
		}
	}
}
