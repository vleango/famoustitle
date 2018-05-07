package models_test

import (
	"github.com/stretchr/testify/suite"
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

func (suite *Suite) TestArticle() {
	article := models.Article{}
	suite.IsType("testing-2134", article.ID)
	suite.IsType("my title", article.Title)
	suite.IsType("my body", article.Body)
	suite.IsType([]string{"tag1", "tag2"}, article.Tags)
	suite.IsType(time.Now(), article.CreatedAt)
	suite.IsType(time.Now(), article.UpdatedAt)
}

func (suite *Suite) TestArticleCreateTitleBlank() {
	article := test.DefaultArticleModel()
	article.Title = ""

	_, err := models.ArticleCreate(article)
	suite.Equal(models.ErrTitleBodyNotProvided, err)
}

func (suite *Suite) TestArticleCreateBodyBlank() {
	article := test.DefaultArticleModel()
	article.Body = ""

	_, err := models.ArticleCreate(article)
	suite.Equal(models.ErrTitleBodyNotProvided, err)
}

func (suite *Suite) TestArticleCreateTagsBlank() {
	article := test.DefaultArticleModel()
	article.Tags = []string{
		"",
	}
	item, err := models.ArticleCreate(article)
	suite.IsType(nil, err)
	suite.Equal(0, len(item.Tags))
}

func (suite *Suite) TestArticleCreateTagsWhitespace() {
	article := test.DefaultArticleModel()
	article.Tags = []string{
		" tag1 ",
		" tag2",
	}
	item, err := models.ArticleCreate(article)
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
	item, err := models.ArticleCreate(article)
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
	item, err := models.ArticleCreate(article)
	suite.IsType(nil, err)
	suite.Equal(2, len(item.Tags))
	suite.Contains(item.Tags, "tag1")
	suite.Contains(item.Tags, "tag2")
}

func (suite *Suite) TestArticleCreateSuccess() {
	article := test.DefaultArticleModel()
	item, err := models.ArticleCreate(article)
	suite.IsType(nil, err)

	suite.Equal(article.Title, item.Title)
	suite.Equal(article.Body, item.Body)
	suite.Equal(article.Tags, item.Tags)
	suite.NotNil(item.ID)
	suite.NotNil(item.CreatedAt)
	suite.NotNil(item.UpdatedAt)
}

func (suite *Suite) TestArticleDestroyRecordFound() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	item, err := models.ArticleDestroy(article)
	suite.IsType(nil, err)
	suite.Equal(article, item)

	_, err = models.ArticleFind(article.ID)
	suite.Equal(models.ErrRecordNotFound, err)
}

func (suite *Suite) TestArticleDestroyRecordNotFound() {
	article := test.DefaultArticleModel()
	article.ID = "not-id"
	_, err := models.ArticleDestroy(article)
	suite.Equal(models.ErrRecordNotFound, err)
}

func (suite *Suite) TestArticleFindAllEmpty() {
	items, err := models.ArticleFindAll()
	suite.IsType(nil, err)
	suite.Equal([]models.Article{}, items)
}

func (suite *Suite) TestArticleFindAllNotEmpty() {
	var articles []models.Article
	article1, _ := models.ArticleCreate(test.DefaultArticleModel())
	article2, _ := models.ArticleCreate(test.DefaultArticleModel())
	articles = append(articles, article1)
	articles = append(articles, article2)

	items, err := models.ArticleFindAll()
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
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	item, err := models.ArticleFind(article.ID)
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
	_, err := models.ArticleFind("not-found-id")
	suite.Equal(models.ErrRecordNotFound, err)
}

func (suite *Suite) TestArticleUpdateSuccess() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	article.Title = "new title"
	article.Body = "new body"

	updatedArticle, err := models.ArticleUpdate(article)
	suite.IsType(nil, err)
	suite.Equal(article.ID, updatedArticle.ID)
	suite.Equal("new title", updatedArticle.Title)
	suite.Equal("new body", updatedArticle.Body)
	suite.Equal(article.CreatedAt.Unix(), updatedArticle.CreatedAt.Unix())
}

func (suite *Suite) TestArticleUpdateSuccessUpdatedAt() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())

	// to change updated_at
	time.Sleep(1 * time.Second)

	updatedArticle, err := models.ArticleUpdate(article)
	suite.IsType(nil, err)
	suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *Suite) TestArticleUpdateTitleBlankBodyPresent() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	originalText := article.Title
	article.Title = ""
	article.Body = "my new body"

	// to change updated_at
	//time.Sleep(1 * time.Second)

	updatedArticle, err := models.ArticleUpdate(article)
	suite.IsType(nil, err)
	suite.Equal(originalText, updatedArticle.Title)
	suite.Equal("my new body", updatedArticle.Body)
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *Suite) TestArticleUpdateBodyBlank() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	originalText := article.Body
	article.Title = "my new title"
	article.Body = ""

	// to change updated_at
	//time.Sleep(1 * time.Second)

	updatedArticle, err := models.ArticleUpdate(article)
	suite.IsType(nil, err)
	suite.Equal("my new title", updatedArticle.Title)
	suite.Equal(originalText, updatedArticle.Body)
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *Suite) TestArticleUpdateTagsPresent() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	article.Tags = []string{
		"tag1",
		"tag2",
		"tag3",
	}

	updatedArticle, err := models.ArticleUpdate(article)
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
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	article.Tags = []string{}

	updatedArticle, err := models.ArticleUpdate(article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(0, len(updatedArticle.Tags))
}

func (suite *Suite) TestArticleUpdateTagsEmptyString() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	article.Tags = []string{
		"",
	}

	updatedArticle, err := models.ArticleUpdate(article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(0, len(updatedArticle.Tags))
}

func (suite *Suite) TestArticleUpdateTagsDup() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	article.Tags = []string{
		"tag1",
		"tag1",
	}

	updatedArticle, err := models.ArticleUpdate(article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(1, len(updatedArticle.Tags))
	suite.Contains(updatedArticle.Tags, "tag1")
}

func (suite *Suite) TestArticleUpdateTagsWhitespace() {
	article, _ := models.ArticleCreate(test.DefaultArticleModel())
	article.Tags = []string{
		" tag1 ",
	}

	updatedArticle, err := models.ArticleUpdate(article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(1, len(updatedArticle.Tags))
	suite.Contains(updatedArticle.Tags, "tag1")
}
