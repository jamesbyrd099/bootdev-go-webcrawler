package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(urlInput string) (string, error) {
	urlData, err := url.Parse(urlInput)
	if err != nil {
		return "", errors.New("error parsing URL: " + urlInput)
	}
	urlHostName := urlData.Hostname()
	urlPath := strings.TrimSuffix(urlData.Path, "/")
	if urlHostName == "" {
		return "", errors.New("error parsing URL: no host name")
	}
	return strings.ToLower(urlHostName + urlPath), nil
}
