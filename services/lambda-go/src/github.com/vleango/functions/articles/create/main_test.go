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
	"github.com/vleango/lib/responses"
	"github.com/vleango/lib/test"
	"log"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	userToken   string
	writerToken string
}

var requestBody RequestArticle

func (suite *Suite) SetupTest() {
	test.CleanDataStores()
	test.CreateArticlesTable()
	tokens := test.CreateUserTable(map[string]interface{}{
		"user": models.User{
			FirstName: "Tha",
			LastName:  "Leang",
			Email:     "tha.leang@test.com",
		},
		"password": "hogehoge",
	}, map[string]interface{}{
		"user": models.User{
			FirstName: "Writer",
			LastName:  "Me",
			Email:     "writer@test.com",
		},
		"password": "hogehoge",
	})

	writer := models.User{
		Email:    "writer@test.com",
		IsWriter: true,
	}

	_, err := dynamodb.UserUpdate(writer)
	if err != nil {
		log.Fatal(err)
	}

	suite.userToken = tokens[0]
	suite.writerToken = tokens[1]

	requestBody = RequestArticle{
		Article: test.DefaultArticleModel(),
	}
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

func (suite *Suite) TestSavingNewRecord() {
	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(jsonRequestBody),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.writerToken),
		},
	}

	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody models.Article
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal("Writer Me", responseBody.Author)
	suite.Equal(requestBody.Article.Title, responseBody.Title)
	suite.Equal(requestBody.Article.Body, responseBody.Body)
	suite.Equal(2, len(responseBody.Tags))
	suite.Contains(responseBody.Tags, "ruby")
	suite.Contains(responseBody.Tags, "rails")
	suite.NotNil(responseBody.ID)
	suite.NotNil(responseBody.CreatedAt)
	suite.NotNil(responseBody.UpdatedAt)

	time.Sleep(2 * time.Second)
	articles, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(1, len(articles))
	suite.Equal("Writer Me", articles[0].Author)
	suite.Equal(requestBody.Article.Title, articles[0].Title)
	suite.Equal(requestBody.Article.Body, articles[0].Body)
	suite.Equal(2, len(articles[0].Tags))
	suite.Contains(articles[0].Tags, "ruby")
	suite.Contains(articles[0].Tags, "rails")
	suite.Equal(responseBody.ID, articles[0].ID)
	suite.Equal(responseBody.CreatedAt, articles[0].CreatedAt)
	suite.Equal(responseBody.UpdatedAt, articles[0].UpdatedAt)
}

func (suite *Suite) TestMissingTags() {
	requestBody.Article.Tags = nil

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(jsonRequestBody),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.writerToken),
		},
	}

	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody models.Article
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal(requestBody.Article.Title, responseBody.Title)
	suite.Equal(requestBody.Article.Body, responseBody.Body)
	suite.Equal([]string{}, responseBody.Tags)
	suite.NotNil(responseBody.ID)
	suite.NotNil(responseBody.CreatedAt)
	suite.NotNil(responseBody.UpdatedAt)
}

func (suite *Suite) TestMissingTitle() {
	requestBody.Article.Title = ""

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(jsonRequestBody),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.writerToken),
		},
	}

	response, err := Handler(context.Background(), request)
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
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.writerToken),
		},
	}

	response, err := Handler(context.Background(), request)
	suite.IsType(nil, err)
	suite.Equal(400, response.StatusCode)

	var responseBody map[string]string
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Equal("missing title and/or body in the HTTP body", responseBody["message"])
}

func (suite *Suite) TestUserNotWriter() {
	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		Body: string(jsonRequestBody),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.userToken),
		},
	}

	response, err := Handler(context.Background(), request)
	var responseBody map[string]string
	json.Unmarshal([]byte(response.Body), &responseBody)

	suite.Nil(err)
	suite.Equal(responses.StatusUnauthorized, response.StatusCode)
	suite.Equal("user not allowed to create articles", responseBody["message"])
}
