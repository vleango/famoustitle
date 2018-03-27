package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/models"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	_, err := models.ArticleDestroy(models.Article{ID: request.PathParameters["id"]})
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
		Body:       string(body),
		StatusCode: 200,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
	}, nil
}

func main() {
	lambda.Start(Handler)
}

type ResponseBody struct {
	Success bool
}
