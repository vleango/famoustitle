package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
)

var (
	DynamoSvc           *dynamodb.DynamoDB
	DynamoArticlesTable = "famoustitle_articles"
	DynamoUsersTable    = "famoustitle_users"

	Session *session.Session
)

func init() {
	var err error
	urlStr := os.Getenv("DYNAMODB_HOST_URL")

	if urlStr == "" {
		Session, err = session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("REGION")),
		})
	} else {
		Session, err = session.NewSession(&aws.Config{
			Region:   aws.String(os.Getenv("REGION")),
			Endpoint: &urlStr,
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	DynamoSvc = dynamodb.New(Session)
}
