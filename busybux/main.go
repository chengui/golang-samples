package main

import "os"

func main() {
	sh := NewShell()
	sh.Run(os.Stdin, os.Stdout)
}
