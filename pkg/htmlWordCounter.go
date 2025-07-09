package htmlWordCounter

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func visit(doc *html.Node, words, pics *int) {
	if doc.Type == html.TextNode {
		for p := doc.Parent; p != nil; p = p.Parent {
			if p.Type == html.ElementNode && (p.Data == "script" || p.Data == "style") {
				return
			}
		}
		*words += len(strings.Fields(doc.Data))
	} else if doc.Type == html.ElementNode && doc.Data == "img" {
		*pics++
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}

func countWordsAndImages(doc *html.Node) (words, pics int) {
	visit(doc, &words, &pics)
	return
}

func Start(raw string) (words, pics int) {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	words, pics = countWordsAndImages(doc)

	return
}
