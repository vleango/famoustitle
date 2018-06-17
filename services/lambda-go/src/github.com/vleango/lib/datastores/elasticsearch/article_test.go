package elasticsearch

import (
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/test"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupTest() {
	test.CleanElasticSearch()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestArticleCreate() {
	item := test.DefaultArticleModel()
	article, err := ArticleCreate(item)
	suite.IsType(nil, err)
	suite.Equal(item.Title, article.Title)
	suite.Equal(item.Body, article.Body)
	suite.Equal(item.Author, article.Author)
	suite.Equal(len(item.Tags), len(article.Tags))
	suite.Contains(article.Tags, "rails")
	suite.Contains(article.Tags, "ruby")
}

func (suite *Suite) TestArticleDestroy() {
	item := test.DefaultArticleModel()
	item.ID = "1234"

	ArticleCreate(item)
	time.Sleep(2 * time.Second)
	articles, _, _ := ArticleFindAll()
	suite.Equal(1, len(articles))

	ArticleDestroy(item)
	time.Sleep(2 * time.Second)
	articles2, _, _ := ArticleFindAll()
	suite.Equal(0, len(articles2))
}

func (suite *Suite) TestArticleDestroyIDNotFound() {
	item := test.DefaultArticleModel()
	_, err := ArticleDestroy(item)
	suite.Equal(ErrSaveFailed.Error(), err.Error())

}

func (suite *Suite) TestArticleFindAll() {
	item1 := test.DefaultArticleModel()
	item1.ID = "1234"
	item2 := test.DefaultArticleModel()
	item2.ID = "abcd"

	ArticleCreate(item1)
	ArticleCreate(item2)
	time.Sleep(2 * time.Second)

	articles, _, _ := ArticleFindAll()
	suite.Equal(2, len(articles))

	for _, article := range articles {
		suite.Equal(item1.Title, article.Title)
		suite.Equal(item1.Body, article.Body)
		suite.Equal(len(item1.Tags), len(article.Tags))
		suite.Contains(article.Tags, "rails")
		suite.Contains(article.Tags, "ruby")

		if article.ID != "1234" {
			suite.Equal(item2.ID, article.ID)
		} else {
			suite.Equal(item1.ID, article.ID)
		}
	}
}

func (suite *Suite) TestArticleFindAllByTag() {
	item1 := test.DefaultArticleModel()
	item1.ID = "1234"
	item2 := test.DefaultArticleModel()
	item2.ID = "abcd"

	ArticleCreate(item1)
	ArticleCreate(item2)
	time.Sleep(2 * time.Second)

	articles, _, _ := ArticleFindAll(map[string]string{
		"tag": "rails",
	})
	suite.Equal(2, len(articles))
}

func (suite *Suite) TestArticleFindAllByTagNotFound() {
	item1 := test.DefaultArticleModel()
	item1.ID = "1234"
	item2 := test.DefaultArticleModel()
	item2.ID = "abcd"

	ArticleCreate(item1)
	ArticleCreate(item2)
	time.Sleep(2 * time.Second)

	articles, _, _ := ArticleFindAll(map[string]string{
		"tag": "web",
	})
	suite.Equal(0, len(articles))
}

func (suite *Suite) TestArticleFindAllByDate() {
	item1 := test.DefaultArticleModel()
	item1.ID = "1234"
	item2 := test.DefaultArticleModel()
	item2.ID = "abcd"

	ArticleCreate(item1)
	ArticleCreate(item2)
	time.Sleep(2 * time.Second)

	articles, _, _ := ArticleFindAll(map[string]string{
		"date": time.Now().Format("2006-01-") + "01",
	})
	suite.Equal(2, len(articles))
}

func (suite *Suite) TestArticleFindAllByDateNotFound() {
	item1 := test.DefaultArticleModel()
	item1.ID = "1234"
	item2 := test.DefaultArticleModel()
	item2.ID = "abcd"

	ArticleCreate(item1)
	ArticleCreate(item2)
	time.Sleep(2 * time.Second)

	articles, _, _ := ArticleFindAll(map[string]string{
		"date": "2000-01-01",
	})
	suite.Equal(0, len(articles))
}

func (suite *Suite) TestArticleFindAllByMatch() {
	item1 := test.DefaultArticleModel()
	item1.ID = "1234"
	item2 := test.DefaultArticleModel()
	item2.ID = "abcd"

	ArticleCreate(item1)
	ArticleCreate(item2)
	time.Sleep(2 * time.Second)

	articles, _, _ := ArticleFindAll(map[string]string{
		"match": "title",
	})
	suite.Equal(2, len(articles))
}

func (suite *Suite) TestArticleFindAllByMatchNotFound() {
	item1 := test.DefaultArticleModel()
	item1.ID = "1234"
	item2 := test.DefaultArticleModel()
	item2.ID = "abcd"

	ArticleCreate(item1)
	ArticleCreate(item2)
	time.Sleep(2 * time.Second)

	articles, _, _ := ArticleFindAll(map[string]string{
		"match": "noooo",
	})
	suite.Equal(0, len(articles))
}

func (suite *Suite) TestArticleUpdate() {
	item := test.DefaultArticleModel()
	item.ID = "1234"

	ArticleCreate(item)
	time.Sleep(2 * time.Second)

	item.Title = "my new title"
	item.Body = "my new body"
	item.Tags = []string{"web"}

	ArticleUpdate(item)
	time.Sleep(2 * time.Second)

	articles, _, _ := ArticleFindAll()
	article := articles[0]
	suite.Equal("my new title", article.Title)
	suite.Equal("my new body", article.Body)
	suite.Equal(1, len(article.Tags))
	suite.Contains(article.Tags, "web")
}
