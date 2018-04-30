package main_test

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	main "github.com/vleango/functions/articles/create"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"testing"
)

type Suite struct {
	suite.Suite
}

var requestBody main.RequestArticle

func (suite *Suite) SetupTest() {
	test.CleanDB()
	test.CreateArticlesTable()

	requestBody = main.RequestArticle{
		Article: test.DefaultArticleModel(),
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestSavingNewRecord() {
	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(jsonRequestBody),
	}

	response, err := main.Handler(request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody models.Article
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal(requestBody.Article.Title, responseBody.Title)
	suite.Equal(requestBody.Article.Body, responseBody.Body)
	suite.Equal(requestBody.Article.Tags, responseBody.Tags)
	suite.NotNil(responseBody.ID)
	suite.NotNil(responseBody.CreatedAt)
	suite.NotNil(responseBody.UpdatedAt)
}

func (suite *Suite) TestMissingTags() {
	requestBody.Article.Tags = nil

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(jsonRequestBody),
	}

	response, err := main.Handler(request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody models.Article
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal(requestBody.Article.Title, responseBody.Title)
	suite.Equal(requestBody.Article.Body, responseBody.Body)
	suite.Nil(responseBody.Tags)
	suite.NotNil(responseBody.ID)
	suite.NotNil(responseBody.CreatedAt)
	suite.NotNil(responseBody.UpdatedAt)
}

func (suite *Suite) TestMissingTitle() {
	requestBody.Article.Title = ""

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(jsonRequestBody),
	}

	response, err := main.Handler(request)
	suite.IsType(nil, err)
	suite.Equal(400, response.StatusCode)

	var responseBody map[string]string
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal("missing title and/or body in the HTTP body", responseBody["message"])
}

func (suite *Suite) TestMissingBody() {
	requestBody.Article.Body = ""

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(jsonRequestBody),
	}

	response, err := main.Handler(request)
	suite.IsType(nil, err)
	suite.Equal(400, response.StatusCode)

	var responseBody map[string]string
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal("missing title and/or body in the HTTP body", responseBody["message"])
}
