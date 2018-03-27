package database

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoSvc *dynamodb.DynamoDB
var defaultRegion = "us-west-2"
var urlStr = "http://db-dynamo:8000"

func init() {

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(defaultRegion),
		Endpoint: &urlStr,
	})

	if err != nil {
		log.Fatal(err)
	}

	DynamoSvc = dynamodb.New(sess)
}
