package main_test

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	"github.com/vleango/functions/articles/update"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"testing"
)

type Suite struct {
	suite.Suite
}

var (
	article models.Article
)

func (suite *Suite) SetupTest() {
	test.CleanDB()
	test.CreateArticlesTable()
	article, _ = models.ArticleCreate(test.DefaultArticleModel())

	// to change updated_at
	//time.Sleep(1 * time.Second)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestUpdateRecordNotFound() {
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

func (suite *Suite) TestUpdateTitle() {
	requestBody := main.RequestArticle{
		Article: article,
	}
	requestBody.Article.Title = "new title"

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Body: string(jsonRequestBody),
	}

	response, err := main.Handler(request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody models.Article
	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.Equal(article.ID, responseBody.ID)
	suite.Equal("new title", responseBody.Title)
	suite.Equal(article.Body, responseBody.Body)
	suite.Equal(article.Tags, responseBody.Tags)

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), responseBody.CreatedAt.Unix())
	//suite.NotEqual(article.UpdatedAt.Unix(), responseBody.UpdatedAt.Unix())
}

func (suite *Suite) TestUpdateBody() {
	requestBody := main.RequestArticle{
		Article: article,
	}
	requestBody.Article.Body = "new body"

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Body: string(jsonRequestBody),
	}

	response, err := main.Handler(request)
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
}

func (suite *Suite) TestUpdateTitleBlankBodyPresent() {
	requestBody := main.RequestArticle{
		Article: article,
	}
	requestBody.Article.Title = ""
	requestBody.Article.Body = "my new body"

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Body: string(jsonRequestBody),
	}

	response, err := main.Handler(request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	updatedArticle, _ := models.ArticleFind(article.ID)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal("my new body", updatedArticle.Body)

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), updatedArticle.CreatedAt.Unix())
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *Suite) TestUpdateTitlePresentBodyBlank() {
	requestBody := main.RequestArticle{
		Article: article,
	}
	requestBody.Article.Title = "my new title"
	requestBody.Article.Body = ""

	jsonRequestBody, _ := json.Marshal(requestBody)
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": article.ID,
		},
		Body: string(jsonRequestBody),
	}

	response, err := main.Handler(request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	updatedArticle, _ := models.ArticleFind(article.ID)
	suite.Equal("my new title", updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), updatedArticle.CreatedAt.Unix())
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}
