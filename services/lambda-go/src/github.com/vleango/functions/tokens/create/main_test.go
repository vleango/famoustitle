package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/auth"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"strings"
	"testing"
)

type Suite struct {
	suite.Suite
	user     models.User
	password string
	token    string
}

func (suite *Suite) SetupTest() {
	test.CleanDataStores()

	suite.password = "hogehoge"
	suite.user = models.User{
		FirstName: "Tha",
		LastName:  "Leang",
		Email:     "tha.leang@test.com",
	}

	tokens := test.CreateUserTable(map[string]interface{}{
		"user":     suite.user,
		"password": suite.password,
	})
	suite.token = tokens[0]
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestMissingEmail() {
	requestBody := map[string]string{
		"password": suite.password,
	}

	b, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(b),
	}

	response, err := Handler(context.Background(), request)

	var responseBody map[string]interface{}
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal(nil, err)
	suite.Equal(400, response.StatusCode)
	suite.Equal(auth.ErrMissingParams.Error(), responseBody["message"])
}

func (suite *Suite) TestMissingPassword() {
	requestBody := map[string]string{
		"email": suite.user.Email,
	}

	b, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(b),
	}

	response, err := Handler(context.Background(), request)

	var responseBody map[string]interface{}
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal(nil, err)
	suite.Equal(400, response.StatusCode)
	suite.Equal(auth.ErrMissingParams.Error(), responseBody["message"])
}

func (suite *Suite) TestEmailDoesNotExist() {
	requestBody := map[string]string{
		"email":    "some-email",
		"password": suite.password,
	}

	b, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(b),
	}

	response, err := Handler(context.Background(), request)

	var responseBody map[string]interface{}
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal(nil, err)
	suite.Equal(400, response.StatusCode)
	suite.Equal("record not found", responseBody["message"])
}

func (suite *Suite) TestPasswordDoesNotMatch() {
	requestBody := map[string]string{
		"email":    suite.user.Email,
		"password": "bad-password",
	}

	b, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(b),
	}

	response, err := Handler(context.Background(), request)

	var responseBody map[string]interface{}
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal(nil, err)
	suite.Equal(400, response.StatusCode)
	suite.Equal("password does not match", responseBody["message"])
}

func (suite *Suite) TestHandler() {
	requestBody := map[string]string{
		"email":    suite.user.Email,
		"password": suite.password,
	}

	b, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(b),
	}

	response, err := Handler(context.Background(), request)

	var responseBody map[string]interface{}
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal(nil, err)
	suite.Equal(200, response.StatusCode)

	// test JWT header and payload (not timestamp because it can differ)
	result1 := strings.Split(suite.token, ".")
	result2 := strings.Split(responseBody["token"].(string), ".")
	suite.Equal(3, len(result1))
	suite.Equal(3, len(result2))
	suite.Equal(result2[0], result1[0])
	suite.Equal(result2[1], result1[1])
}
