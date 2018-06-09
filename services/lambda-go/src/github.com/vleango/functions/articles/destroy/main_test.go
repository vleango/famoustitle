package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	userTokenWithArticle    string
	UserTokenWithoutArticle string
}

var article *models.Article

func (suite *Suite) SetupTest() {
	test.CleanDataStores()
	test.CreateArticlesTable()
	defaultArticle := test.DefaultArticleModel()
	article, _ = dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")

	tokens := test.CreateUserTable(
		map[string]interface{}{
			"user": models.User{
				FirstName: "Tha",
				LastName:  "Leang",
				Email:     "tha.leang@test.com",
				Articles: map[string]string{
					article.ID: article.Title,
				},
			},
			"password": "hogehoge",
		},
		map[string]interface{}{
			"user": models.User{
				FirstName: "Bob",
				LastName:  "Hope",
				Email:     "bob.hope@test.com",
				Articles: map[string]string{
					"not-my-id": "An article that doesn't exist",
				},
			},
			"password": "hogehoge",
		},
	)
	suite.userTokenWithArticle = tokens[0]
	suite.UserTokenWithoutArticle = tokens[1]
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestUnauthenticated() {
	request := events.APIGatewayProxyRequest{}

	response, err := Handler(context.Background(), request)

	var responseBody map[string]interface{}
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal(nil, err)
	suite.Equal(401, response.StatusCode)
	suite.Equal("unauthorized", responseBody["message"])
}

func (suite *Suite) TestNotUserArticle() {
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.UserTokenWithoutArticle),
		},
	}

	response, err := Handler(context.Background(), request)
	var responseBody map[string]interface{}
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal(nil, err)
	suite.Equal(400, response.StatusCode)
	suite.Equal(dynamodb.ErrArticleDoesNotBelong.Error(), responseBody["message"])
}

func (suite *Suite) TestDestroyRecordFound() {
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.userTokenWithArticle),
		},
	}

	elasticsearch.ArticleCreate(*article)
	time.Sleep(1 * time.Second)
	articles, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(1, len(articles))

	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody map[string]bool
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal(true, responseBody["success"])

	time.Sleep(1 * time.Second)
	articles2, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(0, len(articles2))
}

func (suite *Suite) TestDestroyRecordNotFound() {
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "not-my-id",
		},
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.UserTokenWithoutArticle),
		},
	}

	response, err := Handler(context.Background(), request)
	suite.Equal(400, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody map[string]string
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal("record not found", responseBody["message"])
}
