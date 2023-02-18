package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	//for _, link := range visit(nil, doc) {
	//	fmt.Println(link)
	//}
	//fmt.Println(outline(nil, doc))
	//counts := make(map[string]int)
	//visit3(counts, doc)
	//for k, v := range counts {
	//	fmt.Printf("%s\t%d\n", k, v)
	//}
	for _, link := range visit5(nil, doc, "link") {
		fmt.Println(link)
	}
}

func visit1(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit1(links, c)
	}
	return links
}

func visit2(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	if node.NextSibling != nil {
		links = visit2(links, node.NextSibling)
	}
	if node.FirstChild != nil {
		links = visit2(links, node.FirstChild)
	}
	return links
}

func visit3(counts map[string]int, node *html.Node) {
	if node.Type == html.ElementNode {
		counts[node.Data]++
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit3(counts, c)
	}
}

func visit4(texts []string, node *html.Node) []string {
	if node.Type == html.TextNode {
		if node.Data != "script" && node.Data != "style" {
			text := strings.Trim(node.Data, " \t\n")
			if text != "" {
				texts = append(texts, text)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		texts = visit4(texts, c)
	}
	return texts
}

func visit5(links []string, node *html.Node, filter string) []string {
	if node.Type == html.ElementNode && node.Data == filter {
		for _, attr := range node.Attr {
			if attr.Key == "href" || attr.Key == "src" {
				links = append(links, attr.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit5(links, c, filter)
	}
	return links
}

func outline(stack []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
	return stack
}
