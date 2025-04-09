package main

import (
	"URL_shortener/Internal/application"
	"URL_shortener/Internal/domain/interfaces"
	customHttp "URL_shortener/Internal/infrastructure/http"
	"URL_shortener/Internal/infrastructure/persistence"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func handler() {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)

	var urlShorterRepository interfaces.URLShorter = persistence.DynamoShorterRepositoryConstruct(db)
	var urlRedirecterRepository interfaces.URLRedirecter = persistence.DynamoRedirectRepoConstruct(db)

	var urlShortenerUseCase = *application.NewURLShortenerUseCase(urlShorterRepository)
	var urlRedirecter = *application.NewRedirectHandlerUseCase(urlRedirecterRepository)

	shortenHandler := customHttp.NewShortenHandler(urlShortenerUseCase)
	redirectHandler := customHttp.NewRedirectHandler(urlRedirecter)

	http.HandleFunc("/shorten", shortenHandler.ShortenURL)
	http.HandleFunc("/redirect", redirectHandler.RedirectURL)
}

func main() {
	lambda.Start(handler)
}
