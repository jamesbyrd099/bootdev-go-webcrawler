package main

import "fmt"

func printResults(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\nREPORT for %v\n=============================\n", baseURL)
	pairs := sortResults(pages)
	for _, pair := range pairs {
		if pair.key == "" {
			continue
		}
		fmt.Printf("Found %d internal links to %s\n", pair.value, pair.key)
	}
}
