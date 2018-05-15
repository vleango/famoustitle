package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/models"
)

type Response struct {
	Articles []models.Article       `json:"articles"`
	Archives elasticsearch.Archives `json:"archives"`
	Tags     elasticsearch.Tags     `json:"tags"`
}

func main() {
	lambda.Start(Handler)
}

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

	articles, aggregations, err := elasticsearch.ArticleFindAll()
	data := Response{
		Articles: articles,
		Archives: aggregations.Archives,
		Tags:     aggregations.Tags,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
	}, nil
}
