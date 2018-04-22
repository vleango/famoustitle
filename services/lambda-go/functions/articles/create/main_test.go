package main_test

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/vleango/functions/articles/create"
	"testing"
)

func TestHandler(t *testing.T) {

	test := struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		request: events.APIGatewayProxyRequest{Body: "hello"},
		expect:  "error",
		err:     nil,
	}

	response, err := main.Handler(test.request)
	assert.IsType(t, test.err, err)
	assert.Equal(t, test.expect, response.Body)
}
