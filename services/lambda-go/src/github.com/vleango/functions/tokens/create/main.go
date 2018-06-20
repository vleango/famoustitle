package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/auth"
	"github.com/vleango/lib/responses"
	"github.com/vleango/lib/utils"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, _, earlyExit := responses.NewProxyResponse(&ctx, &request, false)
	if earlyExit != nil {
		return *earlyExit, nil
	}

	body := map[string]string{}
	json.Unmarshal([]byte(request.Body), &body)

	user, token, err := auth.GenerateToken(body["email"], body["password"])
	if token == nil || err != nil {
		return response.BadRequest(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	message := map[string]interface{}{
		"token":      *token,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"is_writer":  user.IsWriter,
	}
	b, err := json.Marshal(message)
	if err != nil {
		return response.ServerError(utils.JSONStringWithKey(responses.StatusMsgServerError), err.Error()), nil
	}

	return response.Ok(string(b)), nil
}

func main() {
	lambda.Start(Handler)
}
