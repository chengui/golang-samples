package main

import (
	"fmt"
	"go-poster/poster"
	"os"
)

func main() {
	p := poster.NewPoster()
	path, err := p.Query(os.Args[1])
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println("Path: ", path)
}
