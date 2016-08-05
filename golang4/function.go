package main

import (
	"fmt"
	"math"
	"os"

	"golang.org/x/net/html"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x + y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	fmt.Println(hypot(3, 4))

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinkgs1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	outline(nil, doc)
}
