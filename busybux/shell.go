package main

import (
	"bufio"
	"busybux/builtins"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	Prompt string
}

func NewShell() *Shell {
	return &Shell{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Prompt: "%",
	}
}

func (sh *Shell) Exec(f builtins.Handler, name string, args []string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(sh.Stderr, "%s: %v\n", name, err)
		}
	}()
	if err := f(sh.Stdout, args); err != nil {
		return err
	}
	return nil
}

func (sh *Shell) Run() {
	sc := bufio.NewScanner(sh.Stdin)
	for {
		fmt.Fprintf(sh.Stdout, "%s ", sh.Prompt)
		if !sc.Scan() {
			break
		}
		ss := strings.Fields(strings.TrimSpace(sc.Text()))
		if len(ss) == 0 {
			continue
		}
		name, args := ss[0], ss[1:]
		if f, ok := builtins.Commands[name]; ok {
			if err := sh.Exec(f, name, args); err != nil {
				fmt.Fprintf(sh.Stderr, "%s: %v\n", name, err)
			}
		} else {
			cmd := exec.Command(name, args...)
			cmd.Stdout, cmd.Stderr = sh.Stdout, sh.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(sh.Stderr, "%s: %v\n", name, err)
			}
		}
	}
}
