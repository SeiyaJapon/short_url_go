package persistence

import (
	"URL_shortener/Internal/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

type DynamoRepo struct {
	db *dynamodb.DynamoDB
}

func NewDynamoRepo() *DynamoRepo {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	}))
	db := dynamodb.New(sess)
	return &DynamoRepo{db: db}
}

func (repo *DynamoRepo) Shorten(url *domain.URL) (string, error) {
	shortURL := generateShortURL(url.Id)
	err := repo.SaveURLMapping(url.String(), shortURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (repo *DynamoRepo) SaveURLMapping(originalURL, shortURL string) error {
	item := domain.URLMapping{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling map: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("URLMapping"),
	}

	_, err = repo.db.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return err
	}

	return nil
}

func generateShortURL(id string) string {
	return "http://short.url/" + id
}
