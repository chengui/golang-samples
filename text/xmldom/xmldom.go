package xmldom

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func ParseFromString(source, mimeType string) (*Document, error) {
	doc, err := html.Parse(strings.NewReader(source))
	if err != nil {
		return nil, err
	}
	return NewDocument(doc), nil
}

func OutlineToString(doc *Document) (string, error) {
	var depth int
	buf := &bytes.Buffer{}
	startElement := func(node *html.Node) {
		if node.Type == html.ElementNode {
			buf.WriteString(fmt.Sprintf("%*s<%s>\n", depth*2, "", node.Data))
			depth++

		}
	}
	endElement := func(node *html.Node) {
		if node.Type == html.ElementNode {
			depth--
			buf.WriteString(fmt.Sprintf("%*s</%s>\n", depth*2, "", node.Data))

		}
	}
	var visit func(_ *html.Node, _1, _2 func(*html.Node))
	visit = func(node *html.Node, pre, post func(*html.Node)) {
		if pre != nil {
			pre(node)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			visit(c, pre, post)
		}
		if post != nil {
			post(node)
		}
	}
	visit(doc.Node, startElement, endElement)
	return buf.String(), nil
}

func SerializeToString(doc *Document) (string, error) {
	var depth int
	buf := &bytes.Buffer{}
	startElement := func(node *html.Node) {
		if node.Type == html.ElementNode || node.Type == html.TextNode || node.Type == html.CommentNode {
			if node.Type == html.ElementNode {
				var attrs []string
				for _, attr := range node.Attr {
					attrs = append(attrs, fmt.Sprintf("%s=%q", attr.Key, attr.Val))
				}
				var nodeAttrs string
				if len(attrs) > 0 {
					nodeAttrs += " " + strings.Join(attrs, " ")
				}
				if node.FirstChild == nil {
					buf.WriteString(fmt.Sprintf("%*s<%s%s/>\n", depth*2, "", node.Data, nodeAttrs))
				} else {
					buf.WriteString(fmt.Sprintf("%*s<%s%s>\n", depth*2, "", node.Data, nodeAttrs))
				}
			} else {
				buf.WriteString(fmt.Sprintf("%*s%s\n", depth*2, "", node.Data))
			}
			depth++
		}
	}
	endElement := func(node *html.Node) {
		if node.Type == html.ElementNode || node.Type == html.TextNode || node.Type == html.CommentNode {
			depth--
			if node.Type == html.ElementNode {
				if node.FirstChild != nil {
					buf.WriteString(fmt.Sprintf("%*s</%s>\n", depth*2, "", node.Data))
				}
			}
		}
	}
	var visit func(*html.Node, func(*html.Node), func(*html.Node))
	visit = func(node *html.Node, pre, post func(*html.Node)) {
		if pre != nil {
			pre(node)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			visit(c, pre, post)
		}
		if post != nil {
			post(node)
		}
	}
	visit(doc.Node, startElement, endElement)
	return buf.String(), nil
}

func ExtractLinks(doc *Document) ([]string, error) {
	var visit func([]string, *html.Node, func([]string, *html.Node) []string) []string
	visit = func(result []string, node *html.Node, f func([]string, *html.Node) []string) []string {
		if f != nil {
			result = append(result, f(nil, node)...)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			result = visit(result, c, f)
		}
		return result
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
	return visit(nil, doc.Node, find), nil
}

func ExtractImages(doc *Document) ([]string, error) {
	var visit func([]string, *html.Node, func([]string, *html.Node) []string) []string
	visit = func(result []string, node *html.Node, f func([]string, *html.Node) []string) []string {
		if f != nil {
			result = append(result, f(nil, node)...)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			result = visit(result, c, f)
		}
		return result
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
	return visit(nil, doc.Node, find), nil
}

func ExtractScripts(doc *Document) ([]string, error) {
	var visit func([]string, *html.Node, func([]string, *html.Node) []string) []string
	visit = func(result []string, node *html.Node, f func([]string, *html.Node) []string) []string {
		if f != nil {
			result = append(result, f(nil, node)...)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			result = visit(result, c, f)
		}
		return result
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
	return visit(nil, doc.Node, find), nil
}

func ExtractTexts(doc *Document) ([]string, error) {
	var visit func([]string, *html.Node) []string
	visit = func(texts []string, node *html.Node) []string {
		if node.Type == html.ElementNode && (node.Data == "script" || node.Data == "style") {
			return texts
		}
		if node.Type == html.TextNode {
			text := strings.Trim(node.Data, " \t\n")
			if text != "" {
				texts = append(texts, text)
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			texts = visit(texts, c)
		}
		return texts
	}
	return visit(nil, doc.Node), nil
}
