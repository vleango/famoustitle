package main

import (
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "time"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

var region = "us-west-2"
var tableName = "articles"

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  str := "http://db-dynamo:8000"

  sess, err := session.NewSession(&aws.Config{
    Region: aws.String(region),
    Endpoint: &str,
  })

  if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	svc := dynamodb.New(sess)

  _, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
    TableName: aws.String(tableName),
    Key: map[string]*dynamodb.AttributeValue{
        "id": {
            S: aws.String(request.PathParameters["id"]),
        },
    },
  })

  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

  response := ResponseBody{
    Success: true,
  }

  body, err := json.Marshal(response)
  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

  return events.APIGatewayProxyResponse{
    Body: string(body),
    StatusCode: 200,
    Headers: map[string]string{ "Access-Control-Allow-Origin": "*" },
    }, nil
}

func main() {
	lambda.Start(Handler)
}

type Article struct {
  ID    string `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
  CreatedAt time.Time `json:"created_at"`
}

type ResponseBody struct {
  Success bool
}
