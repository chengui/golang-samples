package dirname

import (
	"flag"
	"fmt"
	"io"
	"path/filepath"
)

var helpFlag bool

func Main(w io.Writer, args []string) error {
	flagSet := flag.NewFlagSet("dirname", flag.ContinueOnError)
	flagSet.BoolVar(&helpFlag, "help", false, "Show this message.")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if helpFlag {
		flagSet.Usage()
		return nil
	}

	return dirname(w, flagSet.Args())
}

func dirname(w io.Writer, args []string) error {
	for _, arg := range args {
		fmt.Fprintln(w, filepath.Dir(arg))
	}
	return nil
}
