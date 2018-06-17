package dynamodb_test // need to change package name to avoid import cycle since this package is 'dynamodb' and lib/test imports 'dynamodb'

import (
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"time"
)

func (suite *ArticleSuite) TestArticleCreateNilArticle() {
	_, err := dynamodb.ArticleCreate(nil, "Tha Leang")
	suite.Equal(dynamodb.ErrTitleBodyNotProvided, err)
}

func (suite *ArticleSuite) TestArticleCreateTitleBlank() {
	article := test.DefaultArticleModel()
	article.Title = ""

	_, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.Equal(dynamodb.ErrTitleBodyNotProvided, err)
}

func (suite *ArticleSuite) TestArticleCreateBodyBlank() {
	article := test.DefaultArticleModel()
	article.Body = ""

	_, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.Equal(dynamodb.ErrTitleBodyNotProvided, err)
}

func (suite *ArticleSuite) TestArticleCreateTagsBlank() {
	article := test.DefaultArticleModel()
	article.Tags = []string{
		"",
	}
	item, err := dynamodb.ArticleCreate(&article, "Tha Leang")
	suite.IsType(nil, err)
	suite.Equal(0, len(item.Tags))
}

func (suite *ArticleSuite) TestArticleCreateTagsWhitespace() {
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

func (suite *ArticleSuite) TestArticleCreateTagsLowerCase() {
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

func (suite *ArticleSuite) TestArticleCreateTagsUnique() {
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

func (suite *ArticleSuite) TestArticleCreateSuccess() {
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

func (suite *ArticleSuite) TestArticleDestroyRecordFound() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	item, err := dynamodb.ArticleDestroy(*article)
	suite.IsType(nil, err)
	suite.Equal(article, item)

	_, err = dynamodb.ArticleFind(article.ID)
	suite.Equal(dynamodb.ErrRecordNotFound, err)
}

func (suite *ArticleSuite) TestArticleDestroyRecordNotFound() {
	article := test.DefaultArticleModel()
	article.ID = "not-id"
	_, err := dynamodb.ArticleDestroy(article)
	suite.Equal(dynamodb.ErrRecordNotFound, err)
}

func (suite *ArticleSuite) TestArticleFindAllEmpty() {
	items, err := dynamodb.ArticleFindAll()
	suite.IsType(nil, err)
	suite.Equal([]models.Article{}, items)
}

func (suite *ArticleSuite) TestArticleFindAllNotEmpty() {
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

func (suite *ArticleSuite) TestArticleFindSuccess() {
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

func (suite *ArticleSuite) TestArticleFindFailure() {
	_, err := dynamodb.ArticleFind("not-found-id")
	suite.Equal(dynamodb.ErrRecordNotFound, err)
}

func (suite *ArticleSuite) TestArticleUpdateSuccess() {
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

func (suite *ArticleSuite) TestArticleUpdateSuccessUpdatedAt() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")

	// to change updated_at
	time.Sleep(2 * time.Second)

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *ArticleSuite) TestArticleUpdateTitleBlankBodyPresent() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	originalText := article.Title
	article.Title = ""
	article.Body = "my new body"

	// to change updated_at
	//time.Sleep(2 * time.Second)

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(originalText, updatedArticle.Title)
	suite.Equal("my new body", updatedArticle.Body)
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *ArticleSuite) TestArticleUpdateBodyBlank() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	originalText := article.Body
	article.Title = "my new title"
	article.Body = ""

	// to change updated_at
	//time.Sleep(2 * time.Second)

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal("my new title", updatedArticle.Title)
	suite.Equal(originalText, updatedArticle.Body)
	//suite.NotEqual(article.UpdatedAt.Unix(), updatedArticle.UpdatedAt.Unix())
}

func (suite *ArticleSuite) TestArticleUpdateTagsPresent() {
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

func (suite *ArticleSuite) TestArticleUpdateTagsEmpty() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	article.Tags = []string{}

	updatedArticle, err := dynamodb.ArticleUpdate(*article)
	suite.IsType(nil, err)
	suite.Equal(article.Title, updatedArticle.Title)
	suite.Equal(article.Body, updatedArticle.Body)
	suite.Equal(0, len(updatedArticle.Tags))
}

func (suite *ArticleSuite) TestArticleUpdateTagsEmptyString() {
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

func (suite *ArticleSuite) TestArticleUpdateTagsDup() {
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

func (suite *ArticleSuite) TestArticleUpdateTagsWhitespace() {
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
