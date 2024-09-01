package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) pageLimitReached() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) >= cfg.maxPages
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	fmt.Printf("adding page visit: %v\n", normalizedURL)
	count, visited := cfg.pages[normalizedURL]
	count++
	cfg.pages[normalizedURL] = count
	return !visited
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()
	if cfg.pageLimitReached() {
		return
	}

	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current: %v\n", rawCurrentURL)
		return
	}
	if cfg.baseURL.Hostname() != parsedCurrent.Hostname() {
		fmt.Printf("skipping host name: %v\n", parsedCurrent.Hostname())
		return
	}
	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing url. url: %v, error: %v\n", rawCurrentURL, err)
		return
	}
	fmt.Printf("crawling url: %v\n", rawCurrentURL)
	isFirst := cfg.addPageVisit(normalizedCurrent)
	if !isFirst {
		fmt.Println("not the first visit...")
		return
	}
	fmt.Printf("making request to %v\n", rawCurrentURL)

	currentHtml, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v\n", err)
		return
	}
	moreUrls, err := getURLsFromHTML(currentHtml, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("error getting urls: %v\n", err)
		return
	}
	for _, u := range moreUrls {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}
}
