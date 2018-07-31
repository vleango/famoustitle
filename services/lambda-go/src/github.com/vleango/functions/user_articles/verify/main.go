package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/responses"
	"github.com/vleango/lib/utils"
)

// verify if token is valid and article belongs to user
// will allow the frontend to know if the current token is allowed to edit an article
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// verify token
	response, user, earlyExit := responses.NewProxyResponse(&ctx, &request, true)
	if earlyExit != nil {
		return *earlyExit, nil
	}

	// get article
	article, err := elasticsearch.ArticleFind(request.PathParameters["id"])
	if err != nil {
		return response.NotFound(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	// verify article belongs to user
	_, err = dynamodb.IsUserArticle(*user, *article)
	if err != nil {
		return response.Unauthorized(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	// return the article
	b, err := json.Marshal(article)
	if err != nil {
		return response.ServerError(utils.JSONStringWithKey(responses.StatusMsgServerError), err.Error()), nil
	}

	return response.Ok(string(b)), nil
}

func main() {
	lambda.Start(Handler)
}
