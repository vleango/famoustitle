package main_test

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	main "github.com/vleango/functions/articles/create"
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

	////////

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{
				Body: "",
			},
			expect:  "Hello Paul",
			err:     nil,
		},
		// {
		// 	// Test that the handler responds ErrNameNotProvided
		// 	// when no name is provided in the HTTP body
		// 	request: events.APIGatewayProxyRequest{Body: ""},
		// 	expect:  "",
		// 	err:     main.ErrNameNotProvided,
		// },
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}
