package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"testing"
)

type Suite struct {
	suite.Suite
	user     models.User
	password string
}

func (suite *Suite) SetupTest() {
	test.CleanDataStores()
	test.CreateUserTable()
	suite.password = "hogehoge"
	suite.user = models.User{
		FirstName: "Tha",
		LastName:  "Leang",
		Email:     "tha.leang@test.com",
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestMissingFirstName() {
	user := suite.user
	user.FirstName = ""
	requestBody := map[string]interface{}{
		"user":                  user,
		"password":              suite.password,
		"password_confirmation": suite.password,
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
	suite.Equal("missing required params", responseBody["message"])
}

func (suite *Suite) TestMissingLastName() {
	user := suite.user
	user.LastName = ""
	requestBody := map[string]interface{}{
		"user":                  user,
		"password":              suite.password,
		"password_confirmation": suite.password,
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
	suite.Equal("missing required params", responseBody["message"])
}

func (suite *Suite) TestMissingEmail() {
	user := suite.user
	user.Email = ""
	requestBody := map[string]interface{}{
		"user":                  user,
		"password":              suite.password,
		"password_confirmation": suite.password,
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
	suite.Equal("missing required params", responseBody["message"])
}

func (suite *Suite) TestMissingPassword() {
	user := suite.user
	requestBody := map[string]interface{}{
		"user":                  user,
		"password_confirmation": suite.password,
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
	suite.Equal("missing required params", responseBody["message"])
}

func (suite *Suite) TestMissingPasswordConfirmation() {
	user := suite.user
	requestBody := map[string]interface{}{
		"user":     user,
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
	suite.Equal("missing required params", responseBody["message"])
}

func (suite *Suite) TestPasswordMismatch() {
	user := suite.user
	user.FirstName = ""
	requestBody := map[string]interface{}{
		"user":                  user,
		"password":              suite.password,
		"password_confirmation": "some-wrong-pass",
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
	suite.Equal("missing required params", responseBody["message"])
}

func (suite *Suite) TestHandler() {
	user := suite.user
	requestBody := map[string]interface{}{
		"user":                  user,
		"password":              suite.password,
		"password_confirmation": suite.password,
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
	suite.NotNil(responseBody["token"])
}
