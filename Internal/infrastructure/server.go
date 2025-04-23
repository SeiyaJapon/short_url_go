package http

import (
	"URL_shortener/Internal/application"
	customHttp "URL_shortener/Internal/infrastructure/http"
	"URL_shortener/Internal/infrastructure/services"
	"URL_shortener/config"
	"net/http"
)

func NewServer() (*http.Server, error) {
	urlShorterRepository, urlRedirecterRepository, err := config.InitializeRepositories()
	if err != nil {
		panic("Failed to initialize repositories: " + err.Error())
	}

	urlShortenerUseCase := *application.NewURLShortenerUseCase(urlShorterRepository)
	urlRedirecter := *application.NewRedirectHandlerUseCase(urlRedirecterRepository)

	methodValidator := &services.HTTPMethodValidator{}
	queryParamProcessor := &services.QueryParamProcessor{}

	shortenController := customHttp.NewShortenController(urlShortenerUseCase)
	redirectController := customHttp.NewRedirectController(urlRedirecter, methodValidator, queryParamProcessor)

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", shortenController.ShortenURL)
	mux.HandleFunc("/redirect", redirectController.RedirectURL)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}, nil
}
