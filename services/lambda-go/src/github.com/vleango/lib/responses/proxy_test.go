package responses

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"os"
	"testing"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupTest() {
	test.CleanDB()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestNewProxyResponseMethodOptions() {
	ctx := context.Background()
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "OPTIONS",
	}
	resp, user, earlyExit := NewProxyResponse(&ctx, &request, false)
	suite.Nil(resp)
	suite.Nil(user)
	suite.Equal(StatusOk, earlyExit.StatusCode)
	suite.Equal("", earlyExit.Body)
}

func (suite *Suite) TestNewProxyResponseUnauthenticated() {
	request := events.APIGatewayProxyRequest{}
	resp, user, earlyExit := NewProxyResponse(nil, &request, true)
	var responseBody map[string]string
	json.Unmarshal([]byte(earlyExit.Body), &responseBody)

	suite.Nil(resp)
	suite.Nil(user)
	suite.Equal(StatusUnauthorized, earlyExit.StatusCode)
	suite.Equal("unauthorized", responseBody["message"])
}

func (suite *Suite) TestNewProxyResponseAuthenticated() {
	tokens := test.CreateUserTable(map[string]interface{}{
		"user": models.User{
			FirstName: "Tha",
			LastName:  "Leang",
			Email:     "tha.leang@test.com",
		},
		"password": "hogehoge",
	})

	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", tokens[0]),
		},
	}

	resp, user, earlyExit := NewProxyResponse(nil, &request, true)
	suite.Equal(&Response{}, resp)
	suite.Nil(earlyExit)
	suite.Equal("Tha", user.FirstName)
	suite.Equal("Leang", user.LastName)
	suite.NotEmpty(user.ID)
}

func (suite *Suite) TestProxyResponseLogOption() {
	response := Response{}
	resp := response.ProxyResponse(100, "hello", []interface{}{"debug msg"})

	defaultHeaders := map[string]string{
		"Access-Control-Allow-Origin":  os.Getenv("DOMAIN_URL"),
		"Access-Control-Allow-Headers": "Content-Type,Authorization",
		"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE",
	}

	suite.Equal(100, resp.StatusCode)
	suite.Equal("hello", resp.Body)
	suite.Equal(defaultHeaders, resp.Headers)
}

func (suite *Suite) TestProxyResponseWithoutOptions() {
	response := Response{}
	resp := response.ProxyResponse(100, "hello", []interface{}{})
	suite.Equal(100, resp.StatusCode)
	suite.Equal("hello", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestOk() {
	response := Response{}
	resp := response.Ok("hi")
	suite.Equal(StatusOk, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestCreated() {
	response := Response{}
	resp := response.Created("hi")
	suite.Equal(StatusCreated, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestAccepted() {
	response := Response{}
	resp := response.Accepted("hi")
	suite.Equal(StatusAccepted, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestNoContent() {
	response := Response{}
	resp := response.NoContent("hi")
	suite.Equal(StatusNoContent, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestBadRequest() {
	response := Response{}
	resp := response.BadRequest("hi")
	suite.Equal(StatusBadRequest, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestUnauthorized() {
	response := Response{}
	resp := response.Unauthorized("hi")
	suite.Equal(StatusUnauthorized, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestForbidden() {
	response := Response{}
	resp := response.Forbidden("hi")
	suite.Equal(StatusForbidden, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestNotFound() {
	response := Response{}
	resp := response.NotFound("hi")
	suite.Equal(StatusNotFound, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestMethodNotAllowed() {
	response := Response{}
	resp := response.MethodNotAllowed("hi")
	suite.Equal(StatusMethodNotAllowed, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestMethodImATeapot() {
	response := Response{}
	resp := response.ImATeapot("hi")
	suite.Equal(StatusImATeapot, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestMethodUnprocessableEntity() {
	response := Response{}
	resp := response.UnprocessableEntity("hi")
	suite.Equal(StatusUnprocessableEntity, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestMethodServerError() {
	response := Response{}
	resp := response.ServerError("hi")
	suite.Equal(StatusServerError, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}

func (suite *Suite) TestServiceUnavailable() {
	response := Response{}
	resp := response.ServiceUnavailable("hi")
	suite.Equal(StatusServiceUnavailable, resp.StatusCode)
	suite.Equal("hi", resp.Body)
	suite.Equal(DefaultHeaders, resp.Headers)
}
