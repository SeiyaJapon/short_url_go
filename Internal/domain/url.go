package domain

import (
	"errors"
	"net/url"
)

type URL struct {
	Id  string
	Url string
}

func NewURL(value string) (*URL, error) {
	if !isValidURL(value) {
		return nil, errors.New("invalid url")
	}
	return &URL{Id: value}, nil
}

func isValidURL(value string) bool {
	_, err := url.ParseRequestURI(value)
	return err == nil
}

func (u *URL) String() string {
	return u.Url
}
