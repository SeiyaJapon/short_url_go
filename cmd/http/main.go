package main

import (
	"URL_shortener/Internal/application"
	"URL_shortener/Internal/domain/interfaces"
	customHttp "URL_shortener/Internal/infrastructure/http"
	"URL_shortener/Internal/infrastructure/persistence"
	"net/http"
)

func main() {
	var urlShorterRepository interfaces.URLShorter = persistence.DynamoShorterRepositoryConstruct()
	var urlRedirecterRepository interfaces.URLRedirecter = '' // TODO: Implement the repository
	var urlShortenerUseCase = *application.NewURLShortenerUseCase(urlShorterRepository)
	var urlRedirecter = *application.NewRedirectHandlerUseCase(urlRedirecterRepository)

	shortenHandler := customHttp.NewShortenHandler(urlShortenerUseCase)
	redirectHandler := customHttp.NewRedirectHandler(urlRedirecter)

	http.HandleFunc("/shorten", shortenHandler.ShortenURL)
	http.HandleFunc("/redirect", redirectHandler.RedirectURL)
}
