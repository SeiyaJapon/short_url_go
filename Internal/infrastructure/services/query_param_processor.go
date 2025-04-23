package services

import (
	"errors"
	"net/url"
	"strings"
)

type QueryParamProcessor struct{}

func (p *QueryParamProcessor) ExtractShortURL(rawQuery string) (string, error) {
	queryParams, err := url.ParseQuery(rawQuery)
	if err != nil {
		return "", errors.New("invalid query parameters")
	}

	shortURL := queryParams.Get("short_url")
	if shortURL == "" {
		return "", errors.New("short_url parameter is missing")
	}

	parts := strings.Split(shortURL, "/")
	return parts[len(parts)-1], nil
}
