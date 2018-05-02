package main_test

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	"github.com/vleango/functions/articles/index"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"testing"
)

type Suite struct {
	suite.Suite
}

var (
	articles     []models.Article
	responseBody main.Response
)

func (suite *Suite) SetupTest() {
	test.CleanDB()
	test.CreateArticlesTable()
	articles = []models.Article{}

	article1, _ := models.ArticleCreate(test.DefaultArticleModel())
	article2, _ := models.ArticleCreate(test.DefaultArticleModel())
	article3, _ := models.ArticleCreate(test.DefaultArticleModel())

	articles = append(articles, article1, article2, article3)

	request := events.APIGatewayProxyRequest{}
	response, err := main.Handler(request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	json.Unmarshal([]byte(response.Body), &responseBody)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestArticles() {
	suite.Equal(len(articles), len(responseBody.Articles))

	// loop through response to check for same values
	for _, responseArticle := range responseBody.Articles {
		for _, createdArticle := range articles {
			if createdArticle.ID == responseArticle.ID {
				suite.Equal(createdArticle.ID, responseArticle.ID)
				suite.Equal(createdArticle.Title, responseArticle.Title)
				suite.Equal(createdArticle.Body, responseArticle.Body)
				suite.Equal(createdArticle.Tags, responseArticle.Tags)

				// convert these to unix epoch to check for matching
				suite.Equal(createdArticle.CreatedAt.Unix(), responseArticle.CreatedAt.Unix())
				suite.Equal(createdArticle.UpdatedAt.Unix(), responseArticle.UpdatedAt.Unix())
			}
		}
	}
}

func (suite *Suite) TestArchives() {
	key := articles[0].CreatedAt.Format("January 2006")
	suite.Equal(len(articles), responseBody.Archives[key])
}

func (suite *Suite) TestTags() {
	defaultTags := []string{"ruby", "rails"}
	suite.Equal(len(defaultTags), len(responseBody.Tags))
	suite.Contains(responseBody.Tags, "ruby")
	suite.Contains(responseBody.Tags, "rails")
}
