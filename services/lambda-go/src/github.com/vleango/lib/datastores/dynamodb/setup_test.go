package dynamodb_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"testing"
)

type ArticleSuite struct {
	suite.Suite
}

type UserSuite struct {
	suite.Suite
	user    models.User
	article models.Article
}

func (suite *ArticleSuite) SetupTest() {
	test.CleanDB()
	test.CreateArticlesTable()
}

func (suite *UserSuite) SetupTest() {
	test.CleanDB()
	test.CreateArticlesTable()

	suite.user = models.User{
		FirstName: "Tha",
		LastName:  "Leang",
		Email:     "tha.leang@test.com",
	}
	test.CreateUserTable(map[string]interface{}{
		"user":     suite.user,
		"password": "hogehoge",
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ArticleSuite))
	suite.Run(t, new(UserSuite))
}
