package main

import (
	"bufio"
	"busybux/builtins"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Shell struct {
	Prompt string
}

func NewShell() *Shell {
	return &Shell{Prompt: "%"}
}

func (sh *Shell) Exec(w io.Writer, f builtins.Handler, name string, args []string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "sh: %v\n", err)
		}
	}()
	if err := f(w, args); err != nil {
		fmt.Fprintf(w, "sh: %v\n", err)
	}
	return nil
}

func (sh *Shell) Run(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)
	for {
		fmt.Fprintf(w, "%s ", sh.Prompt)
		if !sc.Scan() {
			break
		}
		ss := strings.Fields(strings.TrimSpace(sc.Text()))
		if len(ss) == 0 {
			continue
		}
		name, args := ss[0], ss[1:]
		if f, ok := builtins.Commands[name]; ok {
			if err := sh.Exec(w, f, name, args); err != nil {
				fmt.Fprintf(w, "sh: %v\n", err)
			}
		} else {
			cmd := exec.Command(name, args...)
			cmd.Stdout, cmd.Stderr = w, w
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(w, "sh: %v\n", err)
			}
		}
	}
}
