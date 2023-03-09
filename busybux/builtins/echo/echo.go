package echo

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var helpFlag bool
var newlineFlag bool

func Main(w io.Writer, args []string) error {
	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	flagSet.BoolVar(&helpFlag, "help", false, "Show this message.")
	flagSet.BoolVar(&newlineFlag, "n", false, "Do not print the trailing newline character.")
	flagSet.Parse(args)

	if helpFlag {
		flagSet.Usage()
		return nil
	}

	return echo(w, flagSet.Args())
}

func echo(w io.Writer, args []string) error {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + os.ExpandEnv(arg)
		sep = " "
	}
	if newlineFlag {
		fmt.Fprint(w, s)
	} else {
		fmt.Fprintln(w, s)
	}
	return nil
}
