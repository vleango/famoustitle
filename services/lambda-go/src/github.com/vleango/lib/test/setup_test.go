package test

import (
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/models"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupTest() {
	CleanDataStores()
	CreateArticlesTable()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestCleanDataStores() {
	articlesDyanomo, _ := dynamodb.ArticleFindAll()
	suite.Equal(0, len(articlesDyanomo))

	articlesES, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(0, len(articlesES))

	defaultArticle := DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	elasticsearch.ArticleCreate(*article)
	time.Sleep(2 * time.Second)

	articlesDyanomo, _ = dynamodb.ArticleFindAll()
	suite.Equal(1, len(articlesDyanomo))

	articlesES, _, _ = elasticsearch.ArticleFindAll()
	suite.Equal(1, len(articlesES))

	CleanDataStores()

	articlesDyanomo, _ = dynamodb.ArticleFindAll()
	suite.Equal(0, len(articlesDyanomo))

	articlesES, _, _ = elasticsearch.ArticleFindAll()
	suite.Equal(0, len(articlesES))
}

func (suite *Suite) TestCleanDB() {
	articlesDyanomo, _ := dynamodb.ArticleFindAll()
	suite.Equal(0, len(articlesDyanomo))

	defaultArticle := DefaultArticleModel()
	dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")

	articlesDyanomo, _ = dynamodb.ArticleFindAll()
	suite.Equal(1, len(articlesDyanomo))

	CleanDB()

	articlesDyanomo, _ = dynamodb.ArticleFindAll()
	suite.Equal(0, len(articlesDyanomo))
}

func (suite *Suite) TestCleanElasticSerch() {
	articlesES, _, _ := elasticsearch.ArticleFindAll()
	suite.Equal(0, len(articlesES))

	defaultArticle := DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	elasticsearch.ArticleCreate(*article)
	time.Sleep(2 * time.Second)

	articlesES, _, _ = elasticsearch.ArticleFindAll()
	suite.Equal(1, len(articlesES))

	CleanElasticSearch()

	articlesES, _, _ = elasticsearch.ArticleFindAll()
	suite.Equal(0, len(articlesES))
}

func (suite *Suite) TestDefaultArticleModel() {
	model := DefaultArticleModel()
	suite.Equal("Tha", model.Author)
	suite.Equal("this is my title", model.Title)
	suite.Equal("this is my body", model.Body)
	suite.Equal(2, len(model.Tags))
	suite.Contains(model.Tags, "ruby")
	suite.Contains(model.Tags, "rails")
	suite.NotEmpty(model.CreatedAt)
	suite.NotEmpty(model.UpdatedAt)
}

func (suite *Suite) TestCreateUserTable() {
	tokens := CreateUserTable(map[string]interface{}{
		"user": models.User{
			FirstName: "Tha",
			LastName:  "Leang",
			Email:     "tha.leang@test.com",
		},
		"password": "hogehoge",
	})

	user, err := dynamodb.UserFindByEmail("tha.leang@test.com")
	suite.Equal(1, len(tokens))
	suite.Nil(err)
	suite.NotEmpty(user.ID)
	suite.NotEmpty(user.PasswordDigest)
	suite.Equal("Tha", user.FirstName)
	suite.Equal("Leang", user.LastName)
	suite.Equal("tha.leang@test.com", user.Email)
	suite.NotEmpty(user.CreatedAt)
	suite.NotEmpty(user.UpdatedAt)
}
