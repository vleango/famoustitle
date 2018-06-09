package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/responses"
	"github.com/vleango/lib/utils"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, user, earlyExit := responses.NewProxyResponse(&ctx, &request, true)
	if earlyExit != nil {
		return *earlyExit, nil
	}

	item, err := dynamodb.UserArticleDestroy(*user, models.Article{ID: request.PathParameters["id"]})
	if err != nil {
		return response.BadRequest(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	b, err := json.Marshal(map[string]bool{"success": true})
	if err != nil {
		return response.ServerError(utils.JSONStringWithKey(responses.StatusMsgServerError), err.Error()), nil
	}

	elasticsearch.ArticleDestroy(*item)
	return response.Ok(string(b)), nil
}

func main() {
	lambda.Start(Handler)
}
