package cat

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var helpFlag bool
var numberFlag bool

func Main(w io.Writer, args []string) error {
	flagSet := flag.NewFlagSet("cat", flag.ExitOnError)
	flagSet.BoolVar(&helpFlag, "help", false, "Show this message.")
	flagSet.BoolVar(&numberFlag, "n", false, "Number the output lines, starting at 1.")
	flagSet.Parse(args)

	if helpFlag {
		flagSet.Usage()
		return nil
	}

	return cat(w, flagSet.Args())
}

type reader struct {
	io.Reader
	cnt int
}

func (in *reader) Read(p []byte) (int, error) {
	if numberFlag {
		l, err := in.ReadBytes('\n')
		if err != nil && err != io.EOF || len(l) == 0 {
			return copy(p, l), err
		}
		in.cnt++
		s := []byte(fmt.Sprintf("%d ", in.cnt))
		s = append(s, l...)
		return copy(p, s), err
	} else {
		return in.Reader.Read(p)
	}
}

func (in *reader) ReadBytes(delim byte) (s []byte, err error) {
	b := make([]byte, 1)
	for {
		_, err = in.Reader.Read(b)
		if err != nil {
			return
		}
		s = append(s, b...)
		if b[0] == '\n' {
			return
		}
	}
}

func cat(w io.Writer, args []string) error {
	for _, arg := range args {
		f, err := os.Open(os.ExpandEnv(arg))
		if err != nil {
			fmt.Printf("cat: %v\n", err)
			continue
		}
		defer f.Close()
		rd := &reader{f, 0}
		if _, err := io.Copy(w, rd); err != nil {
			return err
		}
	}
	return nil
}
