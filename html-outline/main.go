package main

import (
	"fmt"
	"os"

	"html-outline/outline"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}
	doc, err := outline.Fetch(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(2)
	}
	fmt.Println(outline.Outline(doc))
}
