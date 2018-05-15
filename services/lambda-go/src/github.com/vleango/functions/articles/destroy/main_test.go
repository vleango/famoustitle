package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
}

var article models.Article

func (suite *Suite) SetupTest() {
	test.CleanDataStores()
	test.CreateArticlesTable()

	article, _ = dynamodb.ArticleCreate(test.DefaultArticleModel())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestDestroyRecordFound() {
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
	}

	elasticsearch.ArticleCreate(article)
	time.Sleep(1 * time.Second)
	articles, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(1, len(articles))

	response, err := Handler(request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody ResponseBody
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal(true, responseBody.Success)

	time.Sleep(1 * time.Second)
	articles2, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(0, len(articles2))
}

func (suite *Suite) TestDestroyRecordNotFound() {
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "not-my-id",
		},
	}

	response, err := Handler(request)
	suite.Equal(404, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody map[string]string
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal("record not found", responseBody["message"])
}
