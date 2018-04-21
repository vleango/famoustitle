package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/models"
)

func removeDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v:= range elements {
        encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range encountered {
        result = append(result, key)
    }
    return result
}

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
    Tags: removeDuplicatesUnordered(tags),
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
  Archives map[string]int `json:"archives"`
  Tags []string `json:"tags"`
}
