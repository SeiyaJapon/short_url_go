package persistence

import (
	"URL_shortener/Internal/domain"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoRedirectRepository struct {
	db *dynamodb.DynamoDB
}

func DynamoRedirectRepoConstruct() *DynamoRedirectRepository {
	return &DynamoRedirectRepository{}
}

func (repo *DynamoRedirectRepository) Redirect(url *domain.URL) (string, error) {
	return url.String(), nil
}
