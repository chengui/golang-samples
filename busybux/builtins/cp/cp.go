package cp

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

type Option struct {
	helpFlag     bool
	recusiveFlag bool
	preserveFlag bool
	forceFlag    bool
}

var option = &Option{}

func Main(w io.Writer, args []string) error {
	flagSet := flag.NewFlagSet("cp", flag.ContinueOnError)
	flagSet.BoolVar(&option.helpFlag, "help", false, "Show this message.")
	flagSet.BoolVar(&option.recusiveFlag, "R", false, "Recusive copy")
	flagSet.BoolVar(&option.preserveFlag, "p", false, "preserve the attributes of each source file")
	flagSet.BoolVar(&option.forceFlag, "f", false, "force overwrite")
	flagSet.Usage = func() {
		fmt.Println(w, "cp [-R|-p|-f] <src_file> ...<target>")
		flagSet.PrintDefaults()
	}
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	files := flagSet.Args()
	if option.helpFlag || len(files) < 2 {
		flagSet.Usage()
		return nil
	}

	return cp(w, files)
}

func cp(w io.Writer, files []string) error {
	srcs, dst := files[:len(files)-1], files[len(files)-1]
	for _, src := range srcs {
		fmt.Println(src, dst)
		if isDir(src) {
			if err := copyDirectory(src, dst); err != nil {
				fmt.Fprintln(w, "cp:", err)
			}
		} else {
			if err := copyFile(src, dst); err != nil {
				fmt.Fprintln(w, "cp:", err)
			}
		}
	}
	return nil
}

func copyDirectory(src, dst string) error {
	if err := ensureDir(src, false); err != nil {
		return err
	}
	if err := ensureDir(dst, true); err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcNew := path.Join(src, entry.Name())
		dstNew := path.Join(dst, entry.Name())
		if entry.IsDir() {
			copyDirectory(srcNew, dstNew)
		} else {
			copyFile(srcNew, dstNew)
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	if isDir(dst) {
		dst = path.Join(dst, path.Base(src))
	}
	sf, err := ensureFile(src, false)
	if err != nil {
		return fmt.Errorf("%v %s", err, src)
	}
	defer sf.Close()
	df, err := ensureFile(dst, true)
	if err != nil {
		return fmt.Errorf("%v %s", err, dst)
	}
	defer df.Close()
	if _, err := io.Copy(df, sf); err != nil {
		return fmt.Errorf("copy failed: %v", err)
	}
	return nil
}

func ensureDir(p string, create bool) error {
	if create && isNotExist(p) {
		if err := os.MkdirAll(p, os.ModePerm); err != nil {
			return err
		}
	}
	if isNotExist(p) {
		return fmt.Errorf("%q not exist", p)
	}
	if !isDir(p) {
		return fmt.Errorf("%q not directory", p)
	}
	return nil
}

func ensureFile(p string, create bool) (f *os.File, err error) {
	if create && isNotExist(p) {
		f, err = os.Create(p)
		if err != nil {
			return nil, err
		}
	}
	if isNotExist(p) {
		return nil, fmt.Errorf("%q not exist", p)
	}
	if isDir(p) {
		return nil, fmt.Errorf("%q is directory", p)
	}
	if f == nil {
		f, err = os.Open(p)
	}
	return f, err
}

func isNotExist(p string) bool {
	_, err := os.Stat(p)
	if err != nil && os.IsNotExist(err) {
		return true
	}
	return false
}

func isDir(p string) bool {
	fs, err := os.Stat(p)
	if err != nil {
		return false
	}
	return fs.IsDir()
}
