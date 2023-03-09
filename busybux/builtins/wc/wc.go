package wc

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var opt Option

type Option struct {
	helpFlag  bool
	lineFlag  bool
	wordFlag  bool
	countFlag bool
}

type result struct {
	lines int
	words int
	bytes int
}

func (r *result) Add(t *result) {
	r.lines += t.lines
	r.words += t.words
	r.bytes += t.bytes
}

func (r result) String() string {
	var out []string
	if opt.lineFlag {
		out = append(out, fmt.Sprintf("%8d", r.lines))
	}
	if opt.wordFlag {
		out = append(out, fmt.Sprintf("%8d", r.words))
	}
	if opt.countFlag {
		out = append(out, fmt.Sprintf("%8d", r.bytes))
	}
	if !opt.lineFlag && !opt.wordFlag && !opt.countFlag {
		out = append(out, fmt.Sprintf("%8d", r.lines))
		out = append(out, fmt.Sprintf("%8d", r.words))
		out = append(out, fmt.Sprintf("%8d", r.bytes))
	}
	return strings.Join(out, "")
}

func Main(w io.Writer, args []string) error {
	flagSet := flag.NewFlagSet("wc", flag.PanicOnError)
	flagSet.BoolVar(&opt.helpFlag, "help", false, "show this message.")
	flagSet.BoolVar(&opt.lineFlag, "l", false, "The number of lines")
	flagSet.BoolVar(&opt.wordFlag, "w", false, "The number of words")
	flagSet.BoolVar(&opt.countFlag, "c", false, "the number of bytes")
	flagSet.Parse(args)

	if opt.helpFlag {
		flagSet.Usage()
		return nil
	}

	total := &result{}
	for _, arg := range flagSet.Args() {
		path := os.ExpandEnv(arg)
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(w, "wc: %v\n", err)
			continue
		}
		defer f.Close()

		r, err := wc(w, f, path)
		if err != nil {
			fmt.Fprintf(w, "wc: %v\n", err)
			continue
		}
		total.Add(r)
	}
	if flagSet.NArg() > 1 {
		fmt.Fprintf(w, "%s %s\n", total.String(), "total")
	}
	return nil
}

func wc(w io.Writer, f io.Reader, path string) (*result, error) {
	var r result
	rd := bufio.NewReader(f)
	for {
		l, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return &r, err
		}
		r.lines++
		r.bytes += len(l)
		r.words += len(strings.Fields(l))
	}
	fmt.Fprintf(w, "%s %s\n", r.String(), path)
	return &r, nil
}
