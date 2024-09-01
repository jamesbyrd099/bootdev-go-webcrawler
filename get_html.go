package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)

	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("response body close failed...")
		}
	}(resp.Body)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error level status code: %s", resp.Status)
	}
	if h := resp.Header.Get("Content-Type"); !strings.Contains(h, "text/html") {
		return "", fmt.Errorf("incorrect content type. 'expecting: text/html',recieved: %v", h)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body data: %v", err)
	}
	return string(data), nil
}
