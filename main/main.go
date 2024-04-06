package main

import (
	"Link-parser"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	var htmlparsing string = `<a href="/demos/cyoa/" class="clean-link">
                                <em class="fa fa-star pr"></em>
                                Demo
                            </a>`
	a, err := link.Parse(html.NewTokenizer(strings.NewReader(htmlparsing)))
	if err != nil {
		fmt.Println(err)
		//return
	}
	for i, j := range a {
		fmt.Printf("%d: href: %s text: %s\n", i+1, j.Href, j.Text)
	}
}
