package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url that you want to build a sitemap for")
	flag.Parse()

	fmt.Print(urlFlag)
}

/*
	1. GET the webpage
	2. parse all the links on the webpage
	3. build proper urls with our links
	4. filter out any links that have a diff domain
	5. find all the pages (BFS)
	6. print out XML

*/
