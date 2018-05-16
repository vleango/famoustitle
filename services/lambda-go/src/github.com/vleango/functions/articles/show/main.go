package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/datastores/elasticsearch"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	article, err := elasticsearch.ArticleFind(request.PathParameters["id"])
	if err != nil {
		message := map[string]string{
			"message": err.Error(),
		}
		jsonMessage, _ := json.Marshal(message)

		return events.APIGatewayProxyResponse{
			Body:       string(jsonMessage),
			StatusCode: 404,
		}, nil
	}

	b, err := json.Marshal(article)
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
