package cp

import (
	"os"
	"path"
	"testing"
)

func prepareFile(name, content string) error {
	if isNotExist(path.Dir(name)) {
		os.MkdirAll(path.Dir(name), os.ModePerm)
	}
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(content)
	return nil
}

func isExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func TestCp(t *testing.T) {
	files := []struct {
		name, content string
	}{
		{"tmp/hello.txt", "hello\n"},
		{"tmp/a/hello1.txt", "hello\n"},
		{"tmp/a/hello2.txt", "hello\n"},
		{"tmp/a/b/hello3.txt", "hello\n"},
	}
	os.MkdirAll("tmp/", os.ModePerm)
	defer os.RemoveAll("tmp/")
	for _, file := range files {
		prepareFile(file.name, file.content)
	}

	tests := []struct {
		args []string
		dst  string
		want bool
	}{
		{[]string{"tmp/hello.txt", "tmp/world.txt"}, "tmp/world.txt", true},
		{[]string{"tmp/a/hello1.txt", "tmp/world1.txt"}, "tmp/world1.txt", true},
		{[]string{"tmp/hello.txt", "tmp/404/"}, "tmp/404/hello.txt", false},
		{[]string{"tmp/a/hello2.txt", "tmp/404/"}, "tmp/404/hello2.txt", false},
		{[]string{"tmp/a/", "tmp/b/"}, "tmp/b/b/hello3.txt", true},
		{[]string{"tmp/a/b", "tmp/c/"}, "tmp/c/hello3.txt", true},
		{[]string{"tmp/a/b", "tmp/world.txt"}, "tmp/world.txt/hello3.txt", false},
	}
	for i, test := range tests {
		if err := cp(os.Stdout, test.args); err != nil {
			t.Errorf("error: %v", err)
		}
		if got := isExist(test.dst); got != test.want {
			t.Errorf("#%d %t != %t", i, got, test.want)
		}
	}
}
