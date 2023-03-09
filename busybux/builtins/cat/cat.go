package cat

import (
	"bufio"
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

	for _, arg := range flag.Args() {
		f, err := os.Open(os.ExpandEnv(arg))
		if err != nil {
			fmt.Fprintf(w, "cat: %v\n", err)
			continue
		}
		defer f.Close()
		if err := cat(w, f); err != nil {
			fmt.Fprintf(w, "cat: %v\n", err)
		}
	}
	return nil
}

func cat(w io.Writer, r io.Reader) error {
	if numberFlag {
		cnt := 1
		rd := bufio.NewReader(r)
		for {
			line, err := rd.ReadString('\n')
			if err == io.EOF {
				if len(line) > 0 {
					fmt.Fprintf(w, "%d %s", cnt, line)
				}
				return nil
			}
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "%d %s", cnt, line)
			cnt++
		}
	} else {
		if _, err := io.Copy(w, r); err != nil {
			return err
		}
	}
	return nil
}
