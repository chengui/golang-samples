package h5tool

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func fetch(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func visit(result []string, node *html.Node, f func([]string, *html.Node) []string) []string {
	if f != nil {
		result = append(result, f(nil, node)...)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result = visit(result, c, f)
	}
	return result
}

func FindLinks(url string) ([]string, error) {
	doc, err := fetch(url)
	if err != nil {
		return nil, err
	}
	find := func(links []string, node *html.Node) []string {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		return links
	}
	return visit(nil, doc, find), nil
}

func FindImages(url string) ([]string, error) {
	doc, err := fetch(url)
	if err != nil {
		return nil, err
	}
	find := func(images []string, node *html.Node) []string {
		if node.Type == html.ElementNode && node.Data == "img" {
			for _, attr := range node.Attr {
				if attr.Key == "src" {
					images = append(images, attr.Val)
				}
			}
		}
		return images
	}
	return visit(nil, doc, find), nil
}

func FindScripts(url string) ([]string, error) {
	doc, err := fetch(url)
	if err != nil {
		return nil, err
	}
	find := func(scripts []string, node *html.Node) []string {
		if node.Type == html.ElementNode && node.Data == "script" {
			for _, attr := range node.Attr {
				if attr.Key == "src" {
					scripts = append(scripts, attr.Val)
				}
			}
		}
		return scripts
	}
	return visit(nil, doc, find), nil
}

func FindTexts(url string) ([]string, error) {
	doc, err := fetch(url)
	if err != nil {
		return nil, err
	}
	find := func(texts []string, node *html.Node) []string {
		if node.Type == html.ElementNode && (node.Data == "script" || node.Data == "style") {
			return texts
		}
		if node.Type == html.TextNode {
			text := strings.Trim(node.Data, " \t\n")
			if text != "" {
				texts = append(texts, text)
			}
		}
		return texts
	}
	return visit(nil, doc, find), nil
}
