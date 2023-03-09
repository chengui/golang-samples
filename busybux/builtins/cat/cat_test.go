package cat

import (
	"bytes"
	"os"
	"testing"
)

func prepareFile(fname, content string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(content)
	return nil
}

func TestCat(t *testing.T) {
	files := []struct {
		name, content string
	}{
		{"hello_nl.txt", "hello world\nhello world 2\n"},
		{"hello_nonl.txt", "hello world\nhello world 2"},
	}
	for _, file := range files {
		prepareFile(file.name, file.content)
		defer os.Remove(file.name)
	}

	tests := []struct {
		nflag bool
		args  []string
		want  string
	}{
		{
			true,
			[]string{"hello_nl.txt"},
			"1 hello world\n2 hello world 2\n",
		},
		{
			false,
			[]string{"hello_nl.txt"},
			"hello world\nhello world 2\n",
		},
		{
			true,
			[]string{"hello_nonl.txt"},
			"1 hello world\n2 hello world 2",
		},
		{
			false,
			[]string{"hello_nonl.txt"},
			"hello world\nhello world 2",
		},
	}
	for _, test := range tests {
		var buf bytes.Buffer
		numberFlag = test.nflag
		if err := cat(&buf, test.args); err != nil {
			t.Errorf("error: %v", err)
		}
		if got := buf.String(); got != test.want {
			t.Errorf("%q != %q", got, test.want)
		}
	}
}
