package outline

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func Fetch(url string) (*html.Node, error) {
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

func Outline(node *html.Node) string {
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
	visit(node, startElement, endElement)
	return buf.String()
}

func Format(node *html.Node) string {
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
	visit(node, startElement, endElement)
	return buf.String()
}
