package main

import (
	"fmt"
	"os"

	"html-outline/outline"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}
	doc, err := outline.Fetch(os.Args[2])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(2)
	}
	switch os.Args[1] {
	case "outline":
		fmt.Println(outline.Outline(doc))
	case "full":
		fmt.Println(outline.Full(doc))
	}
}
