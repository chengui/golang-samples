package main

import (
	"fmt"
	"os"

	"go-h5tool/h5tool"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <func> <url>", os.Args[0])
		os.Exit(1)
	}
	switch os.Args[1] {
	case "findlinks":
		links, err := h5tool.FindLinks(os.Args[2])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Println("Find Links Below:")
		for _, link := range links {
			fmt.Println(link)
		}
	case "findimages":
		images, err := h5tool.FindImages(os.Args[2])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Println("Find Images Below:")
		for _, image := range images {
			fmt.Println(image)
		}
	case "findscripts":
		scripts, err := h5tool.FindScripts(os.Args[2])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Println("Find Scripts Below:")
		for _, script := range scripts {
			fmt.Println(script)
		}
	case "findtexts":
		texts, err := h5tool.FindTexts(os.Args[2])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Println("Find Texts Below:")
		for _, text := range texts {
			fmt.Println(text)
		}
	}
}
