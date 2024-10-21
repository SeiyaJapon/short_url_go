package application

import (
	"URL_shortener/Internal/domain"
	"URL_shortener/Internal/domain/interfaces"
)

type URLShortenerUseCase struct {
	urlShorter interfaces.URLShorter
}

func NewURLShortenerUseCase(urlShorter interfaces.URLShorter) *URLShortenerUseCase {
	return &URLShortenerUseCase{urlShorter: urlShorter}
}

func (u *URLShortenerUseCase) ShortenURL(url string) (string, error) {
	newURL, err := domain.NewURL(url)
	if err != nil {
		return "", err
	}

	shortURL, err := u.urlShorter.Shorten(newURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
