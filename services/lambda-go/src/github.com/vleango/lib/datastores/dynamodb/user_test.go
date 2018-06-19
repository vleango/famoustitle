package dynamodb_test

import (
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
)

func (suite *UserSuite) TestUserFindByEmailMissingEmail() {
	user, err := dynamodb.UserFindByEmail("")
	suite.Nil(user)
	suite.Equal(dynamodb.ErrEmailRequired, err)
}

func (suite *UserSuite) TestUserFindByEmailNotFound() {
	user, err := dynamodb.UserFindByEmail("bogus@email.com")
	suite.Nil(user)
	suite.Equal(dynamodb.ErrRecordNotFound, err)
}

func (suite *UserSuite) TestUserFindByEmail() {
	user, err := dynamodb.UserFindByEmail(suite.user.Email)
	suite.Nil(err)
	suite.NotNil(user.ID)
	suite.NotNil(user.PasswordDigest)
	suite.Equal(suite.user.FirstName, user.FirstName)
	suite.Equal(suite.user.LastName, user.LastName)
	suite.Equal(suite.user.Email, user.Email)
}

func (suite *UserSuite) TestUserCreateMissingFirstName() {
	createUser := suite.user
	createUser.FirstName = ""
	user, err := dynamodb.UserCreate(createUser, "hogehoge", "hogehoge")
	suite.Nil(user)
	suite.Equal("missing required params", err.Error())
}

func (suite *UserSuite) TestUserCreateMissingLastName() {
	createUser := suite.user
	createUser.LastName = ""
	user, err := dynamodb.UserCreate(createUser, "hogehoge", "hogehoge")
	suite.Nil(user)
	suite.Equal("missing required params", err.Error())
}

func (suite *UserSuite) TestUserCreateMissingPassword() {
	user, err := dynamodb.UserCreate(suite.user, "", "hogehoge")
	suite.Nil(user)
	suite.Equal("missing required params", err.Error())
}

func (suite *UserSuite) TestUserCreateMissingPasswordConfirmation() {
	user, err := dynamodb.UserCreate(suite.user, "hogehoge", "")
	suite.Nil(user)
	suite.Equal("missing required params", err.Error())
}

func (suite *UserSuite) TestUserCreatePasswordLength() {
	user, err := dynamodb.UserCreate(suite.user, "12345", "12345")
	suite.Nil(user)
	suite.Equal("password min length is 6", err.Error())
}

func (suite *UserSuite) TestUserCreatePasswordMatch() {
	user, err := dynamodb.UserCreate(suite.user, "123456", "abcdef")
	suite.Nil(user)
	suite.Equal("password does not match", err.Error())
}

func (suite *UserSuite) TestUserCreateUniqueEmail() {
	user, err := dynamodb.UserCreate(suite.user, "hogehoge", "hogehoge")
	suite.Nil(user)
	suite.NotNil(err)
}

func (suite *UserSuite) TestUserCreateNames() {
	createUser := suite.user
	createUser.FirstName = "tha"
	createUser.LastName = "leang"
	createUser.Email = "test@example.com"
	user, err := dynamodb.UserCreate(createUser, "hogehoge", "hogehoge")
	suite.Nil(err)
	suite.NotNil(user.ID)
	suite.NotNil(user.PasswordDigest)
	suite.Equal(user.FirstName, "Tha")
	suite.Equal(user.LastName, "Leang")
	suite.Equal(user.Email, "test@example.com")
}

func (suite *UserSuite) TestUserCreate() {
	createUser := suite.user
	createUser.Email = "test@example.com"
	user, err := dynamodb.UserCreate(createUser, "hogehoge", "hogehoge")
	suite.Nil(err)
	suite.NotNil(user.ID)
	suite.NotNil(user.PasswordDigest)
	suite.Equal(user.FirstName, suite.user.FirstName)
	suite.Equal(user.LastName, suite.user.LastName)
	suite.Equal(user.Email, "test@example.com")
}

func (suite *UserSuite) TestUserAddRemoveFromArticleListMissingID() {
	defaultArticle := test.DefaultArticleModel()
	err := dynamodb.UserAddRemoveFromArticleList(suite.user, defaultArticle, true)
	suite.Equal("missing required params", err.Error())
}

func (suite *UserSuite) TestUserAddRemoveFromArticleListMissingTitle() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	article.Title = ""
	err := dynamodb.UserAddRemoveFromArticleList(suite.user, *article, true)
	suite.Equal("missing required params", err.Error())
}

func (suite *UserSuite) TestUserAddRemoveFromArticleListAlreadyExist() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	err := dynamodb.UserAddRemoveFromArticleList(suite.user, *article, true)
	suite.Nil(err)

	user, _ := dynamodb.UserFindByEmail(suite.user.Email)
	suite.Equal(1, len(user.Articles))

	dynamodb.UserAddRemoveFromArticleList(suite.user, *article, true)
	user, _ = dynamodb.UserFindByEmail(suite.user.Email)
	suite.Equal(1, len(user.Articles))
}

func (suite *UserSuite) TestUserAddRemoveFromArticleListDifferentArticles() {
	defaultArticle1 := test.DefaultArticleModel()
	defaultArticle2 := test.DefaultArticleModel()
	article1, _ := dynamodb.ArticleCreate(&defaultArticle1, "Tha Leang")
	article2, _ := dynamodb.ArticleCreate(&defaultArticle2, "Tha Leang")

	user, _ := dynamodb.UserFindByEmail(suite.user.Email)
	err := dynamodb.UserAddRemoveFromArticleList(*user, *article1, true)
	suite.Nil(err)

	user, _ = dynamodb.UserFindByEmail(suite.user.Email)
	err = dynamodb.UserAddRemoveFromArticleList(*user, *article2, true)
	suite.Nil(err)

	user, _ = dynamodb.UserFindByEmail(suite.user.Email)
	suite.Equal(2, len(user.Articles))

	keys := make([]string, 0)
	values := make([]string, 0)
	for key, value := range user.Articles {
		keys = append(keys, key)
		values = append(values, value)
	}

	suite.Contains(keys, article1.ID)
	suite.Contains(keys, article2.ID)
	suite.Contains(values, article1.Title)
	suite.Contains(values, article2.Title)
}

func (suite *UserSuite) TestUserAddRemoveFromArticleListRemoving() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	err := dynamodb.UserAddRemoveFromArticleList(suite.user, *article, true)
	suite.Nil(err)

	user, _ := dynamodb.UserFindByEmail(suite.user.Email)
	suite.Equal(1, len(user.Articles))

	dynamodb.UserAddRemoveFromArticleList(suite.user, *article, false)
	user, _ = dynamodb.UserFindByEmail(suite.user.Email)
	suite.Equal(0, len(user.Articles))
}

func (suite *UserSuite) TestUserArticleDestroyMissingID() {
	article, err := dynamodb.UserArticleDestroy(suite.user, test.DefaultArticleModel())
	suite.Nil(article)
	suite.Equal(dynamodb.ErrArticleDoesNotBelong, err)
}

func (suite *UserSuite) TestUserArticleDestroyDoesNotBelong() {
	article, err := dynamodb.UserArticleDestroy(suite.user, models.Article{ID: "no-id"})
	suite.Nil(article)
	suite.Equal(dynamodb.ErrArticleDoesNotBelong, err)
}

func (suite *UserSuite) TestUserArticleDestroy() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	user, _ := dynamodb.UserFindByEmail(suite.user.Email)
	dynamodb.UserAddRemoveFromArticleList(*user, *article, true)
	user, _ = dynamodb.UserFindByEmail(suite.user.Email)

	deletedArticle, err := dynamodb.UserArticleDestroy(*user, *article)
	suite.Nil(err)
	suite.Equal(article, deletedArticle)

	user, _ = dynamodb.UserFindByEmail(suite.user.Email)
	suite.Equal(0, len(user.Articles))
}

func (suite *UserSuite) TestUserArticleUpdateMissingID() {
	article, err := dynamodb.UserArticleUpdate(suite.user, test.DefaultArticleModel())
	suite.Nil(article)
	suite.Equal(dynamodb.ErrArticleDoesNotBelong, err)
}

func (suite *UserSuite) TestUserArticleUpdateDoesNotBelong() {
	article, err := dynamodb.UserArticleUpdate(suite.user, models.Article{ID: "no-id"})
	suite.Nil(article)
	suite.Equal(dynamodb.ErrArticleDoesNotBelong, err)
}

func (suite *UserSuite) TestUserArticleUpdate() {
	defaultArticle := test.DefaultArticleModel()
	article, _ := dynamodb.ArticleCreate(&defaultArticle, "Tha Leang")
	user, _ := dynamodb.UserFindByEmail(suite.user.Email)
	dynamodb.UserAddRemoveFromArticleList(*user, *article, true)
	user, _ = dynamodb.UserFindByEmail(suite.user.Email)

	article.Title = "my new title"
	article.Body = "my new body"
	article.Tags = []string{"new", "tag"}

	updatedArticle, err := dynamodb.UserArticleUpdate(*user, *article)
	suite.Nil(err)
	suite.Equal("my new title", updatedArticle.Title)
	suite.Equal("my new body", updatedArticle.Body)
	suite.Equal(2, len(updatedArticle.Tags))
	suite.Contains(defaultArticle.Tags, "new")
	suite.Contains(defaultArticle.Tags, "tag")

	user, _ = dynamodb.UserFindByEmail(suite.user.Email)
	suite.Equal(1, len(user.Articles))
	suite.Equal(updatedArticle.Title, user.Articles[updatedArticle.ID])
}
