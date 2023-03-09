package commands

import (
	"busybux/commands/dup"
	"busybux/commands/echo"
	"busybux/commands/wc"
	"io"
)

var Commands map[string]Handler

type Handler func(io.Writer, []string) error

func init() {
	Commands = map[string]Handler{
		"dup":  dup.Main,
		"echo": echo.Main,
		"wc":   wc.Main,
	}

}
