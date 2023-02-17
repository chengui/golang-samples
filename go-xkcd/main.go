package main

import (
	"fmt"
	"go-xkcd/xkcd"
	"os"
)

func main() {
	x := xkcd.NewXkcd()
	x.Load(500, 600)
	fmt.Println("Title: ", os.Args[1])
	fmt.Println("Link: ", x.Find(os.Args[1]))
}
