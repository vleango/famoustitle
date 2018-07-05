package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jinzhu/now"
	"github.com/stretchr/testify/suite"
	"github.com/vleango/lib/datastores/elasticsearch"
	"github.com/vleango/lib/test"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupTest() {
	test.CleanDataStores()
	test.CreateArticlesTable()

	article := test.DefaultArticleModel()
	elasticsearch.ArticleCreate(article)
	elasticsearch.ArticleCreate(article)
	time.Sleep(2 * time.Second)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestHandler() {
	request := events.APIGatewayProxyRequest{}
	response, err := Handler(context.Background(), request)
	suite.Equal(200, response.StatusCode)
	suite.IsType(nil, err)

	var responseBody Response
	json.Unmarshal([]byte(response.Body), &responseBody)
	t, _ := time.Parse(time.RFC3339, responseBody.Archives.Buckets[0].KeyAsString)

	suite.Equal(1, len(responseBody.Archives.Buckets))
	suite.Equal(2, responseBody.Archives.Buckets[0].DocCount)
	suite.Equal(now.BeginningOfMonth().Unix(), t.Unix())
	suite.Equal(now.BeginningOfMonth().Unix()*1000, int64(responseBody.Archives.Buckets[0].Key.(float64)))
}
