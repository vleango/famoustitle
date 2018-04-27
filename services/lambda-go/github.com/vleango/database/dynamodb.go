package database

import (
	"log"
  "os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoSvc *dynamodb.DynamoDB
var defaultRegion = "us-west-2"
var urlStr = "http://db-dynamo:8000"

func init() {
  switch os.Getenv("APP_ENV") {
  case "TEST":
    urlStr = "http://db-dynamo-test:8000"
  case "CI":
    urlStr = "http://localhost:8000"
  }

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(defaultRegion),
		Endpoint: &urlStr,
	})

	if err != nil {
		log.Fatal(err)
	}

	DynamoSvc = dynamodb.New(sess)
}
