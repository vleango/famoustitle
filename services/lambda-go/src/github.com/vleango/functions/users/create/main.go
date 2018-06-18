package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vleango/lib/auth"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/responses"
	"github.com/vleango/lib/utils"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, _, earlyExit := responses.NewProxyResponse(&ctx, &request, false)
	if earlyExit != nil {
		return *earlyExit, nil
	}

	requestUser := RequestUser{}
	json.Unmarshal([]byte(request.Body), &requestUser)

	// create user
	item, err := dynamodb.UserCreate(requestUser.User, requestUser.Password, requestUser.PasswordConfirmation)
	if err != nil {
		return response.BadRequest(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	// generate token
	user, token, err := auth.GenerateToken(item.Email, requestUser.Password)
	if token == nil || err != nil {
		return response.BadRequest(utils.JSONStringWithKey(err.Error()), err.Error()), nil
	}

	message := map[string]string{
		"token":      *token,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
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

type RequestUser struct {
	User                 models.User `json:"user"`
	Password             string      `json:"password"`
	PasswordConfirmation string      `json:"password_confirmation"`
}
