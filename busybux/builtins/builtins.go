package builtins

import (
	"busybux/builtins/basename"
	"busybux/builtins/cat"
	"busybux/builtins/cp"
	"busybux/builtins/dirname"
	"busybux/builtins/dup"
	"busybux/builtins/echo"
	"busybux/builtins/wc"
	"busybux/builtins/wget"
	"io"
)

var Commands map[string]Handler

type Handler func(io.Writer, []string) error

func init() {
	Commands = map[string]Handler{
		"basename": basename.Main,
		"dirname":  dirname.Main,
		"echo":     echo.Main,
		"cat":      cat.Main,
		"dup":      dup.Main,
		"wc":       wc.Main,
		"cp":       cp.Main,
		"wget":     wget.Main,
	}

}
