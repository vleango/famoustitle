package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/responses"
	"github.com/vleango/lib/utils"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, _, earlyExit := responses.NewProxyResponse(&ctx, &request, false)
	if earlyExit != nil {
		return *earlyExit, nil
	}

	article, err := elasticsearch.ArticleFind(request.PathParameters["id"])
	if err != nil {
		return response.NotFound(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	if !article.Published {
		return response.Unauthorized(utils.JSONStringWithKey("unauthorized"), "unauthorized"), nil
	}

	b, err := json.Marshal(article)
	if err != nil {
		return response.ServerError(utils.JSONStringWithKey(responses.StatusMsgServerError), err.Error()), nil
	}

	return response.Ok(string(b)), nil
}

func main() {
	lambda.Start(Handler)
}
