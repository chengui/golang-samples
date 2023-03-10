package basename

import (
	"flag"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type Option struct {
	helpFlag   bool
	allFlag    bool
	suffixFlag string
}

var option = &Option{}

func Main(w io.Writer, args []string) error {
	flagSet := flag.NewFlagSet("basename", flag.ContinueOnError)
	flagSet.BoolVar(&option.helpFlag, "help", false, "Show this message.")
	flagSet.BoolVar(&option.allFlag, "a", false, "every argument is treated as a string")
	flagSet.StringVar(&option.suffixFlag, "s", "", "strip suffix")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if option.helpFlag {
		flagSet.Usage()
		return nil
	}

	return basename(w, flagSet.Args())
}

func basename(w io.Writer, args []string) error {
	var name, suffix string
	if !option.allFlag && option.suffixFlag == "" && len(args) == 2 {
		suffix, args = args[1], args[:1]
	} else {
		suffix = option.suffixFlag
	}
	for _, arg := range args {
		name = strings.TrimSuffix(filepath.Base(arg), suffix)
		fmt.Fprintln(w, name)
	}
	return nil
}
