package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="..."> )in an htmk document
type Link struct {
	href string
	text string
}

// Parse will take a HTML document and return
// a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}
