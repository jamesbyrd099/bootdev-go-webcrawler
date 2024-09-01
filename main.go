package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	var rawBaseUrl string
	maxConcurrency := 1
	maxPages := 50
	// cliArgs := os.Args
	switch len(os.Args) {
	case 1:
		fmt.Println("no website provided")
		os.Exit(1)
	case 2:
		rawBaseUrl = os.Args[1]
		fmt.Println("max concurrency and max pages not provided. using defaults...")
		break
	case 3:
		rawBaseUrl = os.Args[1]
		mc, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("invalid max concurrency input: %v\n", err)
			os.Exit(1)
		}

		maxConcurrency = mc
		break
	case 4:
		rawBaseUrl = os.Args[1]
		mc, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("invalid max concurrency input: %v\n", err)
			os.Exit(1)
		}
		maxConcurrency = mc
		mp, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("invalid page count input: %v\n", err)
			os.Exit(1)
		}
		maxPages = mp
		break
	default:
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", rawBaseUrl)
	pages := make(map[string]int)
	baseURL, err := url.Parse(rawBaseUrl)
	if err != nil {
		fmt.Println("Could not parse base url...")
		return
	}
	mu := sync.Mutex{}
	concurrencyControl := make(chan struct{}, maxConcurrency)
	wg := sync.WaitGroup{}

	cfg := config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 &mu,
		concurrencyControl: concurrencyControl,
		wg:                 &wg,
		maxPages:           maxPages,
	}

	//tStart := time.Now()
	wg.Add(1)
	go cfg.crawlPage(rawBaseUrl)
	wg.Wait()
	//tEnd := time.Now()

	fmt.Println("\nfinished crawl: results below:")
	fmt.Printf("Max concurrency: %d, Max pages: %d\n\n", maxConcurrency, maxPages)
	//crawlCount := 0
	//for key, value := range pages {
	//	fmt.Printf("key: %s, count: %d\n", key, value)
	//	crawlCount++
	//}
	//
	//fmt.Printf("\n%v urls visited in %v\n", crawlCount, tEnd.Sub(tStart))
	printResults(pages, rawBaseUrl)
}
