package dynamodb_test // need to change package name to avoid import cycle since this package is 'dynamodb' and lib/test imports 'dynamodb'

import (
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/datastores/dynamodb"
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

func (suite *Suite) TestArticleCreateTitleBlank() {
	article := test.DefaultArticleModel()
	article.Title = ""

	_, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.Equal(dynamodb.ErrTitleBodyNotProvided, err)
}

func (suite *Suite) TestArticleCreateBodyBlank() {
	article := test.DefaultArticleModel()
	article.Body = ""

	_, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.Equal(dynamodb.ErrTitleBodyNotProvided, err)
}

func (suite *Suite) TestArticleCreateTagsBlank() {
	article := test.DefaultArticleModel()
	article.Tags = []string{
		"",
	}
	item, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.IsType(nil, err)
	suite.Equal(0, len(item.Tags))
}

func (suite *Suite) TestArticleCreateTagsWhitespace() {
	article := test.DefaultArticleModel()
	article.Tags = []string{
		" tag1 ",
		" tag2",
	}
	item, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.IsType(nil, err)
	suite.Equal(2, len(item.Tags))
	suite.Contains(item.Tags, "tag1")
	suite.Contains(item.Tags, "tag2")
}

func (suite *Suite) TestArticleCreateTagsLowerCase() {
	article := test.DefaultArticleModel()
	article.Tags = []string{
		"TAG1",
		"tag2",
	}
	item, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.IsType(nil, err)
	suite.Equal(2, len(item.Tags))
	suite.Contains(item.Tags, "tag1")
	suite.Contains(item.Tags, "tag2")
}

func (suite *Suite) TestArticleCreateTagsUnique() {
	article := test.DefaultArticleModel()
	article.Tags = []string{
		"tag1",
		"tag2",
		"tag1",
	}
	item, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.IsType(nil, err)
	suite.Equal(2, len(item.Tags))
	suite.Contains(item.Tags, "tag1")
	suite.Contains(item.Tags, "tag2")
}

func (suite *Suite) TestArticleCreateSuccess() {
	article := test.DefaultArticleModel()
	item, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.IsType(nil, err)

	suite.Equal(article.Title, item.Title)
	suite.Equal(article.Body, item.Body)
	suite.Equal(len(article.Tags), len(item.Tags))
	suite.Contains(item.Tags, "ruby")
	suite.Contains(item.Tags, "rails")

	suite.NotNil(item.ID)
	suite.NotNil(item.CreatedAt)
	suite.NotNil(item.UpdatedAt)
}

func (suite *Suite) TestArticleDestroyRecordFound() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	item, err := dynamodb.ArticleDestroy(*article)
	suite.IsType(nil, err)
	suite.Equal(article, item)

	_, err = dynamodb.ArticleFind(article.ID)
	suite.Equal(dynamodb.ErrRecordNotFound, err)
}

func (suite *Suite) TestArticleDestroyRecordNotFound() {
	article := test.DefaultArticleModel()
	article.ID = "not-id"
	_, err := dynamodb.ArticleDestroy(article)
	suite.Equal(dynamodb.ErrRecordNotFound, err)
}

func (suite *Suite) TestArticleFindAllEmpty() {
	items, err := dynamodb.ArticleFindAll()
	suite.IsType(nil, err)
	suite.Equal([]models.Article{}, items)
}

func (suite *Suite) TestArticleFindAllNotEmpty() {
	var articles []models.Article
	author := "Tha Leang"
	defaultArticle := test.DefaultArticleModel()
	article1, _ := dynamodb.ArticleCreate(&defaultArticle, author)
	article2, _ := dynamodb.ArticleCreate(&defaultArticle, author)
	articles = append(articles, *article1)
	articles = append(articles, *article2)

	items, err := dynamodb.ArticleFindAll()
	suite.IsType(nil, err)
	suite.Equal(2, len(items))

	// loop through response to check for same values
	for _, fetchedArticle := range items {
		for _, createdArticle := range articles {
			if createdArticle.ID == fetchedArticle.ID {
				suite.Equal(createdArticle.ID, fetchedArticle.ID)
				suite.Equal(createdArticle.Title, fetchedArticle.Title)
				suite.Equal(createdArticle.Body, fetchedArticle.Body)
				suite.Equal(createdArticle.Tags, fetchedArticle.Tags)

				// convert these to unix epoch to check for matching
				suite.Equal(createdArticle.CreatedAt.Unix(), fetchedArticle.CreatedAt.Unix())
				suite.Equal(createdArticle.UpdatedAt.Unix(), fetchedArticle.UpdatedAt.Unix())
			}
		}
	}
}

func (suite *Suite) TestArticleFindSuccess() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	item, err := dynamodb.ArticleFind(article.ID)
	suite.IsType(nil, err)
	suite.Equal(article.ID, item.ID)
	suite.Equal(article.Title, item.Title)
	suite.Equal(article.Body, item.Body)
	suite.Equal(article.Tags, item.Tags)

	// convert these to unix epoch to check for matching
	suite.Equal(article.CreatedAt.Unix(), item.CreatedAt.Unix())
	suite.Equal(article.UpdatedAt.Unix(), item.UpdatedAt.Unix())
}

func (suite *Suite) TestArticleFindFailure() {
	_, err := dynamodb.ArticleFind("not-found-id")
	suite.Equal(dynamodb.ErrRecordNotFound, err)
}

func (suite *Suite) TestArticleUpdateSuccess() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	article.Title = "new title"
	article.Body = "new body"

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(article.ID, updatedArticle.ID)
	suite.Equal("new title", updatedArticle.Title)
	suite.Equal("new body", updatedArticle.Body)
	suite.Equal(article.CreatedAt.Unix(), updatedArticle.CreatedAt.Unix())
}

func (suite *Suite) TestArticleUpdateSuccessUpdatedAt() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")

	// to change updated_at
	time.Sleep(1 * time.Second)

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *Suite) TestArticleUpdateTitleBlankBodyPresent() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	originalText := article.Title
	article.Title = ""
	article.Body = "my new body"

	// to change updated_at
	//time.Sleep(1 * time.Second)

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(originalText, updatedArticle.Title)
	suite.Equal("my new body", updatedArticle.Body)
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *Suite) TestArticleUpdateBodyBlank() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	originalText := article.Body
	article.Title = "my new title"
	article.Body = ""

	// to change updated_at
	//time.Sleep(1 * time.Second)

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal("my new title", updatedArticle.Title)
	suite.Equal(originalText, updatedArticle.Body)
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *Suite) TestArticleUpdateTagsPresent() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	article.Tags = []string{
		"tag1",
		"tag2",
		"tag3",
	}

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
	suite.Equal(3, len(updatedArticle.Tags))
	suite.Contains(updatedArticle.Tags, "tag1")
	suite.Contains(updatedArticle.Tags, "tag2")
	suite.Contains(updatedArticle.Tags, "tag3")
}

func (suite *Suite) TestArticleUpdateTagsEmpty() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	article.Tags = []string{}

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(0, len(updatedArticle.Tags))
}

func (suite *Suite) TestArticleUpdateTagsEmptyString() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	article.Tags = []string{
		"",
	}

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(0, len(updatedArticle.Tags))
}

func (suite *Suite) TestArticleUpdateTagsDup() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	article.Tags = []string{
		"tag1",
		"tag1",
	}

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(1, len(updatedArticle.Tags))
	suite.Contains(updatedArticle.Tags, "tag1")
}

func (suite *Suite) TestArticleUpdateTagsWhitespace() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	article.Tags = []string{
		" tag1 ",
	}

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(1, len(updatedArticle.Tags))
	suite.Contains(updatedArticle.Tags, "tag1")
}
