package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
)

var DynamoSvc *dynamodb.DynamoDB
var defaultRegion = "us-west-2"

var urlStr = ""

func init() {
	switch os.Getenv("APP_ENV") {
	case "development":
		urlStr = "http://db-dynamo:8000"
	case "test":
		urlStr = "http://db-dynamo-test:8000"
	case "ci":
		urlStr = "http://localhost:8000"
	case "production":
		urlStr = ""
	}

	var sess *session.Session
	var err error

	if urlStr == "" {
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String(defaultRegion),
		})
	} else {
		sess, err = session.NewSession(&aws.Config{
			Region:   aws.String(defaultRegion),
			Endpoint: &urlStr,
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	DynamoSvc = dynamodb.New(sess)
}
