package main_test

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	main "github.com/vleango/functions/articles/show"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupTest() {
	test.CleanDB()
	test.CreateArticlesTable()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestShowRecordFound() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
	}
	response, err := main.Handler(request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody map[string]interface{}
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal(article.ID, responseBody["id"])
	suite.Equal(article.Title, responseBody["title"])
	suite.Equal(article.Body, responseBody["body"])

	suite.Equal(len(article.Tags), len(responseBody["tags"].([]interface{})))
	suite.Contains(responseBody["tags"], "ruby")
	suite.Contains(responseBody["tags"], "rails")

	// convert these to unix epoch to check for matching
	responseCreatedAt, _ := time.Parse(time.RFC3339, responseBody["created_at"].(string))
	responseUpdatedAt, _ := time.Parse(time.RFC3339, responseBody["updated_at"].(string))
	suite.Equal(article.CreatedAt.Unix(), responseCreatedAt.Unix())
	suite.Equal(article.UpdatedAt.Unix(), responseUpdatedAt.Unix())
}

func (suite *Suite) TestShowRecordNotFound() {
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "not-found-id",
		},
	}
	response, err := main.Handler(request)
	suite.Equal(404, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody map[string]string
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal("record not found", responseBody["message"])
}
