package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"problems_go/link"
	"strings"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url that you want to build a sitemap for")
	flag.Parse()

	fmt.Print(*urlFlag)

	pages := get(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL

	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))

}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

/*
	1. GET the webpage
	2. parse all the links on the webpage
	3. build proper urls with our links
	4. filter out any links that have a diff domain
	5. find all the pages (BFS)
	6. print out XML

*/
