package basename

import (
	"bytes"
	"testing"
)

func TestBasename(t *testing.T) {
	tests := []struct {
		option *Option
		args   []string
		want   string
	}{
		{
			&Option{false, true, ""},
			[]string{"/a/b/cde"},
			"cde\n",
		},
		{
			&Option{false, false, ""},
			[]string{"/a/b/cde"},
			"cde\n",
		},
		{
			&Option{false, true, ""},
			[]string{"/a/b/cde", "de"},
			"cde\nde\n",
		},
		{
			&Option{false, false, ""},
			[]string{"/a/b/cde", "de"},
			"c\n",
		},
		{
			&Option{false, true, ""},
			[]string{"/a/b/cde", "de", "de"},
			"cde\nde\nde\n",
		},
		{
			&Option{false, false, ""},
			[]string{"/a/b/cde", "de", "de"},
			"cde\nde\nde\n",
		},
		{
			&Option{false, true, "e"},
			[]string{"/a/b/cde", "de"},
			"cd\nd\n",
		},
		{
			&Option{false, false, "e"},
			[]string{"/a/b/cde", "de"},
			"cd\nd\n",
		},
		{
			&Option{false, true, "e"},
			[]string{"/a/b/cde", "de", "de"},
			"cd\nd\nd\n",
		},
		{
			&Option{false, false, "e"},
			[]string{"/a/b/cde", "de", "de"},
			"cd\nd\nd\n",
		},
	}
	for i, test := range tests {
		var buf bytes.Buffer
		option = test.option
		if err := basename(&buf, test.args); err != nil {
			t.Errorf("error: %v", err)
		}
		if got := buf.String(); got != test.want {
			t.Errorf("#%d %q != %q", i, got, test.want)
		}
	}
}
