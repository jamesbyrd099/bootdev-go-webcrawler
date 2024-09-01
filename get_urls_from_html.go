package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

// type Node html.Node

func getRelativeUrls(n *html.Node, currentUrls *[]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				href, err := url.Parse(a.Val)
				if err != nil {
					fmt.Printf("failed to parse href: %v -> %v", href, err)
					continue
				}
				*currentUrls = append(*currentUrls, a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getRelativeUrls(c, currentUrls)
	}
}

func getURLsFromHTML(htmlBody, rawBaseUrl string) ([]string, error) {
	base, err := url.Parse(rawBaseUrl)
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(htmlBody)
	root, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	// fmt.Println(root.String())
	var relativeUrls []string
	getRelativeUrls(root, &relativeUrls)
	var urls []string
	for _, relativeUrl := range relativeUrls {
		href, _ := url.Parse(relativeUrl)
		urls = append(urls, base.ResolveReference(href).String())
	}
	return urls, nil
}
