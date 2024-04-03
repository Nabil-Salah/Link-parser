package main

import (
	"Link-parser"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	var htmlparsing string = `<html>
<head>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
<h1>Social stuffs</h1>
<div>
    <a href="https://www.twitter.com/joncalhoun">
        Check me out on twitter
        <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
        Gophercises is on <strong>Github</strong>!
    </a>
</div>
</body>
</html>`
	a, err := link.Parse(html.NewTokenizer(strings.NewReader(htmlparsing)))
	if err != nil {
		fmt.Println(err)
		//return
	}
	for i, j := range a {
		fmt.Printf("%d: href: %s text: %s\n", i+1, j.Href, j.Text)
	}
}
