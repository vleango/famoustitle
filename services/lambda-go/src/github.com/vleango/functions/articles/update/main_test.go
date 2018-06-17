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

var (
	article *models.Article
)

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

	// to change updated_at
	//time.Sleep(2 * time.Second)
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

func (suite *Suite) TestUpdateRecordNotFound() {
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "not-found-id",
		},
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.userTokenWithArticle),
		},
	}
	response, err := Handler(context.Background(), request)
	suite.Equal(404, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody map[string]string
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal("record not found", responseBody["message"])
}

func (suite *Suite) TestUpdateTitle() {
	requestBody := RequestArticle{
		Article: *article,
	}
	requestBody.Article.Title = "new title"

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Body: string(jsonRequestBody),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.userTokenWithArticle),
		},
	}

	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody models.Article
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal(article.ID, responseBody.ID)
	suite.Equal("new title", responseBody.Title)
	suite.Equal(article.Body, responseBody.Body)

	suite.Equal(len(article.Tags), len(responseBody.Tags))
	suite.Contains(responseBody.Tags, "ruby")
	suite.Contains(responseBody.Tags, "rails")

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), responseBody.CreatedAt.Unix())
	//suite.NotEqual(article.UpdatedAt.Unix(), responseBody.UpdatedAt.Unix())

	time.Sleep(2 * time.Second)
	articles, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(article.ID, articles[0].ID)
	suite.Equal("new title", articles[0].Title)
	suite.Equal(article.Body, articles[0].Body)

	suite.Equal(len(article.Tags), len(articles[0].Tags))
	suite.Contains(articles[0].Tags, "ruby")
	suite.Contains(articles[0].Tags, "rails")

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), articles[0].CreatedAt.Unix())
	suite.Equal(responseBody.UpdatedAt.Unix(), articles[0].UpdatedAt.Unix())
}

func (suite *Suite) TestUpdateBody() {
	requestBody := RequestArticle{
		Article: *article,
	}
	requestBody.Article.Body = "new body"

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Body: string(jsonRequestBody),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.userTokenWithArticle),
		},
	}

	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody models.Article
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal(article.ID, responseBody.ID)
	suite.Equal(article.Title, responseBody.Title)
	suite.Equal("new body", responseBody.Body)

	suite.Equal(len(article.Tags), len(responseBody.Tags))
	suite.Contains(responseBody.Tags, "ruby")
	suite.Contains(responseBody.Tags, "rails")

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), responseBody.CreatedAt.Unix())
	//suite.NotEqual(article.UpdatedAt.Unix(), responseBody.UpdatedAt.Unix())

	time.Sleep(2 * time.Second)
	articles, _, _ := elasticsearch.ArticleFindAll()

	suite.Equal(article.ID, articles[0].ID)
	suite.Equal(article.Title, articles[0].Title)
	suite.Equal("new body", articles[0].Body)

	suite.Equal(len(article.Tags), len(articles[0].Tags))
	suite.Contains(articles[0].Tags, "ruby")
	suite.Contains(articles[0].Tags, "rails")

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), articles[0].CreatedAt.Unix())
	suite.Equal(responseBody.UpdatedAt.Unix(), articles[0].UpdatedAt.Unix())
}

func (suite *Suite) TestUpdateTitleBlankBodyPresent() {
	requestBody := RequestArticle{
		Article: *article,
	}
	requestBody.Article.Title = ""
	requestBody.Article.Body = "my new body"

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Body: string(jsonRequestBody),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.userTokenWithArticle),
		},
	}

	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	updatedArticle, _ := dynamodb.ArticleFind(article.ID)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal("my new body", updatedArticle.Body)

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), updatedArticle.CreatedAt.Unix())
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())

	time.Sleep(2 * time.Second)
	articles, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(article.ID, articles[0].ID)
	suite.Equal(article.Title, articles[0].Title)
	suite.Equal("my new body", articles[0].Body)

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), articles[0].CreatedAt.Unix())
}

func (suite *Suite) TestUpdateTitlePresentBodyBlank() {
	requestBody := RequestArticle{
		Article: *article,
	}
	requestBody.Article.Title = "my new title"
	requestBody.Article.Body = ""

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Body: string(jsonRequestBody),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", suite.userTokenWithArticle),
		},
	}

	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	updatedArticle, _ := dynamodb.ArticleFind(article.ID)
	suite.Equal("my new title", updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), updatedArticle.CreatedAt.Unix())
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())

	time.Sleep(2 * time.Second)
	articles, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(article.ID, articles[0].ID)
	suite.Equal("my new title", articles[0].Title)
	suite.Equal(article.Body, articles[0].Body)

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), articles[0].CreatedAt.Unix())
}

// TODO need tag updates
