package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/models"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// only needed in dev? since template.yml doesn't work??
	if request.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key",
				"Access-Control-Allow-Methods": "GET, POST, OPTIONS, PUT, DELETE",
			},
		}, nil
	}

	// check if article exist
	_, err := models.ArticleFind(request.PathParameters["id"])
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

	requestArticle := RequestArticle{
		Article: models.Article{
			ID: request.PathParameters["id"],
		},
	}
	json.Unmarshal([]byte(request.Body), &requestArticle)
	item, err := models.ArticleUpdate(requestArticle.Article)
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
