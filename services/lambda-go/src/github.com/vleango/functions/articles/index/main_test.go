package main

import (
	"context"
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

type EmptySuite struct {
	suite.Suite
}

var (
	articles     []models.Article
	responseBody Response
)

func (suite *Suite) SetupTest() {
	test.CleanDataStores()
	test.CreateArticlesTable()
	articles = []models.Article{}
	author := "Tha Leang"
	defaultArticle := test.DefaultArticleModel()

	article1, _ := dynamodb.ArticleCreate(&defaultArticle, author)
	elasticsearch.ArticleCreate(*article1)

	article2, _ := dynamodb.ArticleCreate(&defaultArticle, author)
	elasticsearch.ArticleCreate(*article2)

	article3, _ := dynamodb.ArticleCreate(&defaultArticle, author)
	elasticsearch.ArticleCreate(*article3)

	articles = append(articles, *article1, *article2, *article3)
	time.Sleep(2 * time.Second)

	request := events.APIGatewayProxyRequest{}
	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	json.Unmarshal([]byte(response.Body), &responseBody)
}

func (suite *EmptySuite) SetupTest() {
	test.CleanDataStores()
	test.CreateArticlesTable()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
	suite.Run(t, new(EmptySuite))
}

func (suite *Suite) TestArticles() {
	for _, article := range articles {
		elasticsearch.ArticleCreate(article)
	}

	time.Sleep(2 * time.Second)
	esArticles, _, _ := elasticsearch.ArticleFindAll()

	suite.Equal(len(articles), len(esArticles))
	suite.Equal(len(articles), len(responseBody.Articles))

	// loop through response to check for same values
	for _, responseArticle := range responseBody.Articles {
		for _, createdArticle := range articles {
			if createdArticle.ID == responseArticle.ID {
				suite.Equal(createdArticle.ID, responseArticle.ID)
				suite.Equal(createdArticle.Title, responseArticle.Title)
				suite.Nil(responseArticle.Body)
				suite.Equal(createdArticle.Tags, responseArticle.Tags)

				// convert these to unix epoch to check for matching
				suite.Equal(createdArticle.CreatedAt.Unix(), responseArticle.CreatedAt.Unix())
				suite.Equal(createdArticle.UpdatedAt.Unix(), responseArticle.UpdatedAt.Unix())
			}
		}
	}

	// loop through response to check for same values for ES
	for _, esArticle := range esArticles {
		for _, createdArticle := range articles {
			if createdArticle.ID == esArticle.ID {
				suite.Equal(createdArticle.ID, esArticle.ID)
				suite.Equal(createdArticle.Title, esArticle.Title)
				suite.Equal(createdArticle.Body, esArticle.Body)
				suite.Equal(createdArticle.Tags, esArticle.Tags)

				// convert these to unix epoch to check for matching
				suite.Equal(createdArticle.CreatedAt.Unix(), esArticle.CreatedAt.Unix())
				suite.Equal(createdArticle.UpdatedAt.Unix(), esArticle.UpdatedAt.Unix())
			}
		}
	}
}

func (suite *Suite) TestTags() {
	defaultTags := []string{"ruby", "rails"}

	suite.Equal(len(defaultTags), len(responseBody.Tags.Buckets))
	for _, bucket := range responseBody.Tags.Buckets {
		suite.Contains(defaultTags, bucket.Key)
	}
}

func (suite *EmptySuite) TestEmptyArticles() {
	request := events.APIGatewayProxyRequest{}
	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	json.Unmarshal([]byte(response.Body), &responseBody)
	suite.NotNil(responseBody.Articles)
	suite.NotNil(responseBody.Tags.Buckets)
}
