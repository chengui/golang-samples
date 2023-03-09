package xmldom

import (
	"golang.org/x/net/html"
)

type Document struct {
	doctype string
	*html.Node
}

func NewDocument(doc *html.Node) *Document {
	return &Document{
		doctype: "",
		Node:    doc,
	}
}

func (d *Document) CreateElement(tagName string) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: tagName,
	}
}

func (d *Document) CreateTextNode(data string) *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: data,
	}
}

func (d *Document) CreateComment(data string) *html.Node {
	return &html.Node{
		Type: html.CommentNode,
		Data: data,
	}
}

func (d *Document) CreateAttribute(name string) *html.Attribute {
	return &html.Attribute{
		Key: name,
	}
}

func (d *Document) GetElementsByTagName(tagName []string) []*html.Node {
	contains := func(tags []string, tag string) bool {
		for _, t := range tags {
			if t == tag {
				return true
			}
		}
		return false
	}
	var nodes []*html.Node
	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if contains(tagName, n.Data) {
				nodes = append(nodes, n)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(d.Node)
	return nodes
}

func (d *Document) GetElementById(id string) *html.Node {
	var current *html.Node
	startElement := func(node *html.Node) bool {
		if node.Type == html.ElementNode {
			current = node
			for _, attr := range node.Attr {
				if attr.Key == "id" && attr.Val == id {
					return false
				}
			}
		}
		return true
	}
	var visit func(*html.Node, func(*html.Node) bool)
	visit = func(node *html.Node, pre func(*html.Node) bool) {
		if pre != nil {
			if !pre(node) {
				return
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			visit(c, pre)
		}
	}
	visit(d.Node, startElement)
	return current
}
