package xmldom

import "golang.org/x/net/html"

type Element struct {
	*html.Node
	tagName string
}

func NewElement(node *html.Node) *Element {
	return &Element{node, node.Data}
}

func (e *Element) GetElementByTagName(name []string) []*html.Node {
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
			if contains(name, n.Data) {
				nodes = append(nodes, n)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(e.Node)
	return nodes
}

func (d *Element) GetElementById(id string) *html.Node {
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
