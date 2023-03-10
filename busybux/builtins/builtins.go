package builtins

import (
	"busybux/builtins/basename"
	"busybux/builtins/cat"
	"busybux/builtins/dirname"
	"busybux/builtins/dup"
	"busybux/builtins/echo"
	"busybux/builtins/wc"
	"io"
)

var Commands map[string]Handler

type Handler func(io.Writer, []string) error

func init() {
	Commands = map[string]Handler{
		"echo":     echo.Main,
		"cat":      cat.Main,
		"dup":      dup.Main,
		"wc":       wc.Main,
		"basename": basename.Main,
		"dirname":  dirname.Main,
	}

}
