package persistence

import (
	"URL_shortener/Internal/domain"
	"encoding/hex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"math/rand"
)

type DynamoShorterRepository struct {
	db *dynamodb.DynamoDB
}

func DynamoShorterRepositoryConstruct(db *dynamodb.DynamoDB) *DynamoShorterRepository {
	return &DynamoShorterRepository{db: db}
}

func (repo *DynamoShorterRepository) Shorten(url *domain.URL) (string, error) {
	shortCode := generateShortCode(url.Url)
	shortURL := "http://short.url/" + shortCode

	err := repo.saveURLMapping(url.Url, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (repo *DynamoShorterRepository) saveURLMapping(originalURL, shortURL string) error {
	item := domain.URLMapping{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("URLMapping"),
	}

	_, err = repo.db.PutItem(input)
	return err
}

func generateShortCode(url string) string {
	uniqueID, err := generateUniqueId()
	if err != nil {
		return ""
	}
	return uniqueID[:8]
}

func generateUniqueId() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
