package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/models"
)

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

	requestArticle := RequestArticle{}
	json.Unmarshal([]byte(request.Body), &requestArticle)

	item, err := dynamodb.ArticleCreate(requestArticle.Article)
	if err != nil {
		message := map[string]string{
			"message": err.Error(),
		}
		jsonMessage, _ := json.Marshal(message)

		return events.APIGatewayProxyResponse{
			Body:       string(jsonMessage),
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type",
			},
		}, nil
	}

	b, err := json.Marshal(item)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	elasticsearch.ArticleCreate(item)
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
	Article models.Article `json:"article"`
}
