package main

import (
    "fmt"
    "cgo-buffer/wrapper"
)

func main() {
    buff := buffer.NewBuffer(1024)
    defer buff.Delete()

    s := "hello world\x00"
    copy(buff.Data(), []byte(s))
    buff.Print()
    fmt.Println("size=", buff.Size())

    t := " again\x00"
    buff.Append(t)
    buff.Print()
    fmt.Println("size=", buff.Size())
}
