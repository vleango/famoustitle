package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/utils"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	articles, err := models.ArticleFindAll()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	archives := make(map[string]int)
	tags := []string{}
	for _, article := range articles {
		key := article.CreatedAt.Format("January 2006")
		archives[key] += 1

		for _, tag := range article.Tags {
			tags = append(tags, tag)
		}
	}

	data := Response{
		Articles: articles,
		Archives: archives,
		Tags:     utils.RemoveStringDuplicatesUnordered(tags),
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

func main() {
	lambda.Start(Handler)
}

type Response struct {
	Articles []models.Article `json:"articles"`
	Archives map[string]int   `json:"archives"`
	Tags     []string         `json:"tags"`
}
