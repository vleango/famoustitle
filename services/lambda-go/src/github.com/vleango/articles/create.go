package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
)

var region = "us-west-2"
var tableName = "articles"

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if request.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type",
			},
		}, nil
	}

	str := "http://db-dynamo:8000"

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: &str,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	requestArticle := RequestArticle{}
	json.Unmarshal([]byte(request.Body), &requestArticle)

	// TODO need error message
	if requestArticle.Article.Title == "" || requestArticle.Article.Body == "" {
		return events.APIGatewayProxyResponse{
			Body:       "title and/or body is blank",
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type",
			},
		}, nil
	}

	svc := dynamodb.New(sess)

	item := Article{
		ID:        fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil)),
		Title:     requestArticle.Article.Title,
		Body:      requestArticle.Article.Body,
		CreatedAt: time.Now(),
	}

	av, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)

	b, err := json.Marshal(item)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
	}, nil
}

func main() {
	lambda.Start(Handler)
}

type RequestArticle struct {
	Article Article `json:"article"`
}

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
