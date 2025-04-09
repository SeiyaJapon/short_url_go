package persistence

import (
	"URL_shortener/Internal/domain"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoRedirectRepository struct {
	db *dynamodb.DynamoDB
}

func DynamoRedirectRepoConstruct(db *dynamodb.DynamoDB) *DynamoRedirectRepository {
	return &DynamoRedirectRepository{db: db}
}

func (repo *DynamoRedirectRepository) Redirect(url *domain.URL) (string, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("URLMapping"),
		Key: map[string]*dynamodb.AttributeValue{
			"ShortURL": {
				S: aws.String(url.Url),
			},
		},
	}

	result, err := repo.db.GetItem(input)
	if err != nil {
		return "", err
	}

	if result.Item == nil {
		return "", errors.New("URL not found")
	}

	originalURL := result.Item["OriginalURL"].S
	return *originalURL, nil
}
