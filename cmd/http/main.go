package main

import (
	"URL_shortener/Internal/application"
	"URL_shortener/Internal/domain/interfaces"
	customHttp "URL_shortener/Internal/infrastructure/http"
	"URL_shortener/Internal/infrastructure/persistence"
	"net/http"
)

func main() {
	var urlRepo interfaces.URLShorter = persistence.NewDynamoRepo()
	var urlShortenerUseCase = *application.NewURLShortenerUseCase(urlRepo)

	handler := customHttp.NewHandler(urlShortenerUseCase)

	http.HandleFunc("/shorten", handler.ShortenURL)
}
