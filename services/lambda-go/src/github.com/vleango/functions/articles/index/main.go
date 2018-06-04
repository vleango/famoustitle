package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/responses"
	"github.com/vleango/lib/utils"
)

type Response struct {
	Articles []models.Article   `json:"articles"`
	Tags     elasticsearch.Tags `json:"tags"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, _, earlyExit := responses.NewProxyResponse(&ctx, &request, false)
	if earlyExit != nil {
		return *earlyExit, nil
	}

	articles, aggregations, err := elasticsearch.ArticleFindAll(request.QueryStringParameters)
	if err != nil {
		return response.ServerError(utils.JSONStringWithKey(responses.StatusMsgServerError), err.Error()), nil
	}

	data := Response{
		Articles: articles,
		Tags:     aggregations.Tags,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return response.ServerError(utils.JSONStringWithKey(responses.StatusMsgServerError), err.Error()), nil
	}

	return response.Ok(string(b)), nil
}
