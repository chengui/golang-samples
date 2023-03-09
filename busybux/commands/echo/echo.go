package echo

import (
	"fmt"
	"io"
)

func Main(w io.Writer, args []string) error {
	echo1(w, args)
	return nil
}

func echo1(w io.Writer, args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func echo2(w io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}
