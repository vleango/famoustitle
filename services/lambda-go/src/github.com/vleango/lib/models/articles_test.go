package models

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupTest() {
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestArticle() {
	article := Article{}
	suite.IsType("testing-2134", article.ID)
	suite.IsType("my title", article.Title)
	suite.IsType("my body", article.Body)
	suite.IsType([]string{"tag1", "tag2"}, article.Tags)
	suite.IsType(time.Now(), article.CreatedAt)
	suite.IsType(time.Now(), article.UpdatedAt)
}
