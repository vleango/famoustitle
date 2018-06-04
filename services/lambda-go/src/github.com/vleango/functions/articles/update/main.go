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
	response, _, earlyExit := responses.NewProxyResponse(&ctx, &request, true)
	if earlyExit != nil {
		return *earlyExit, nil
	}

	// check if article exist
	_, err := dynamodb.ArticleFind(request.PathParameters["id"])
	if err != nil {
		return response.NotFound(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	requestArticle := RequestArticle{
		Article: models.Article{
			ID: request.PathParameters["id"],
		},
	}
	json.Unmarshal([]byte(request.Body), &requestArticle)
	item, err := dynamodb.ArticleUpdate(requestArticle.Article)
	if err != nil {
		return response.BadRequest(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	b, err := json.Marshal(item)
	if err != nil {
		return response.ServerError(utils.JSONStringWithKey(responses.StatusMsgServerError), err.Error()), nil
	}

	elasticsearch.ArticleUpdate(item)
	return response.Ok(string(b)), nil
}

func main() {
	lambda.Start(Handler)
}

type RequestArticle struct {
	Article models.Article `json:"article"`
}
