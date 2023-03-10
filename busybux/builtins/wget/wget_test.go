package wget

import (
	"bytes"
	"os"
	"testing"
)

func TestWget(t *testing.T) {
	tests := []struct {
		option  *Option
		urls    []string
		outfile string
	}{
		{
			&Option{false, true, 30, "", "Wget/0.0.1"},
			[]string{"http://www.baidu.com"},
			"index.html",
		},
	}
	for i, test := range tests {
		var buf bytes.Buffer
		option = test.option
		if err := wget(&buf, test.urls); err != nil {
			t.Errorf("error :%v", err)
		}
		t.Log(buf.String())
		if b, err := os.ReadFile(test.outfile); err != nil || len(b) == 0 {
			t.Errorf("#%d: %v, %d", i, err, len(b))
		} else {
			os.Remove(test.outfile)
		}
	}
}
