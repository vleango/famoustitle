package main

import (
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "time"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

var region = "us-west-2"
var tableName = "articles"

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  // only needed in dev? since template.yml doesn't work??
  if request.HTTPMethod == "OPTIONS" {
    return events.APIGatewayProxyResponse{
      Body: "",
      StatusCode: 200,
      Headers: map[string]string{
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key",
        "Access-Control-Allow-Methods": "GET, POST, OPTIONS, PUT, DELETE",
      },
    }, nil
  }

  str := "http://db-dynamo:8000"

  sess, err := session.NewSession(&aws.Config{
    Region: aws.String(region),
    Endpoint: &str,
  })

  if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

  requestArticle := RequestArticle{}
  json.Unmarshal([]byte(request.Body), &requestArticle)

  // TODO need error message
  if requestArticle.Article.Title == "" && requestArticle.Article.Body == "" {
    return events.APIGatewayProxyResponse{
      Body: "title and/or body is blank",
      StatusCode: 401,
      Headers: map[string]string{
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Headers": "Content-Type",
      },
    }, nil
  }

	svc := dynamodb.New(sess)

  updateExpression := []string{}
  attributeValue := map[string]*dynamodb.AttributeValue{}

  if requestArticle.Article.Title != "" {
    updateExpression = append(updateExpression, "title = :title")
    attributeValue[":title"] = &dynamodb.AttributeValue{S: aws.String(requestArticle.Article.Title)}
  }

  if requestArticle.Article.Body != "" {
    updateExpression = append(updateExpression, "body = :body")
    attributeValue[":body"] = &dynamodb.AttributeValue{S: aws.String(requestArticle.Article.Body)}
  }

  input := &dynamodb.UpdateItemInput{
    ExpressionAttributeValues: attributeValue,
    Key: map[string]*dynamodb.AttributeValue{
        "id": {
            S: aws.String(request.PathParameters["id"]),
        },
    },
    ReturnValues:     aws.String("UPDATED_NEW"),
    TableName: aws.String(tableName),
    UpdateExpression: aws.String("set " + strings.Join(updateExpression, ", ")),
  }

  item, err := svc.UpdateItem(input)

  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

  b, err := json.Marshal(item)
  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }

  return events.APIGatewayProxyResponse{
    Body: string(b),
    StatusCode: 200,
    Headers: map[string]string{ "Access-Control-Allow-Origin": "*" },
    }, nil
}

func main() {
	lambda.Start(Handler)
}

type RequestArticle struct {
  Article Article `json:"article"`
}

type Article struct {
  ID    string `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
  CreatedAt time.Time `json:"created_at"`
}
