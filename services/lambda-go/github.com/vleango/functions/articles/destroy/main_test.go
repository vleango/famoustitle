package main_test

import (
  "encoding/json"
  "testing"
	"github.com/aws/aws-lambda-go/events"
  "github.com/stretchr/testify/suite"
	main "github.com/vleango/functions/articles/destroy"
  "github.com/vleango/lib/models"
  "github.com/vleango/lib/test"
)

type Suite struct {
    suite.Suite
}

var article models.Article

func (suite *Suite) SetupTest() {
  test.CleanDB()
  test.CreateArticlesTable()

  article, _ = models.ArticleCreate(test.DefaultArticleModel())
}

func TestSuite(t *testing.T) {
    suite.Run(t, new(Suite))
}

func (suite *Suite) TestDestroyRecordFound() {
  request := events.APIGatewayProxyRequest{
    PathParameters: map[string]string{
      "id": article.ID,
    },
  }

  response, err := main.Handler(request)
  suite.Equal(200, response.StatusCode)
  suite.IsType(nil, err)

  var responseBody main.ResponseBody
  json.Unmarshal([]byte(response.Body), &responseBody)
  suite.Equal(true, responseBody.Success)
}

func (suite *Suite) TestDestroyRecordNotFound() {
  request := events.APIGatewayProxyRequest{
    PathParameters: map[string]string{
      "id": "not-my-id",
    },
  }

  response, err := main.Handler(request)
  suite.Equal(404, response.StatusCode)
  suite.IsType(nil, err)

  var responseBody map[string]string
  json.Unmarshal([]byte(response.Body), &responseBody)
  suite.Equal("record not found", responseBody["message"])
}
