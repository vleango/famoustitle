package responses

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/vleango/lib/auth"
	"github.com/vleango/lib/logs"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/utils"
)

var (
	DefaultHeaders = map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Content-Type",
	}

	StatusOk                  = 200
	StatusCreated             = 201
	StatusAccepted            = 202
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusImATeapot           = 418
	StatusUnprocessableEntity = 422
	StatusTooManyRequests     = 429
	StatusServerError         = 500
	StatusServiceUnavailable  = 503

	StatusMsgServerError = "internal server error"
)

// Response type
type Response struct {
	Headers map[string]string
}

// NewProxyResponse returns Response
func NewProxyResponse(ctx *context.Context, request *events.APIGatewayProxyRequest, authenticate bool) (resp *Response, user *models.User, earlyExit *events.APIGatewayProxyResponse) {
	var err error

	if request.HTTPMethod == "OPTIONS" {
		rsp := Response{}
		proxyResponse := rsp.Ok("")
		return nil, nil, &proxyResponse
	}

	if authenticate {
		user, err = auth.AuthenticateUser(request.Headers["Authorization"])
		if err != nil {
			rsp := Response{}
			proxyResponse := rsp.Unauthorized(
				utils.JSONStringWithKey(err.Error()),
				fmt.Sprintf("unauthorized: %v", request.Headers["Authorization"]),
			)
			return nil, nil, &proxyResponse
		}
	}

	return &Response{}, user, nil
}

// ProxyResponse returns an APIGateway Proxy Response
// options[0] is a log message
func (r *Response) ProxyResponse(status int, body string, options []interface{}) events.APIGatewayProxyResponse {

	if len(options) > 0 {
		logs.DebugMessage(status, options[0].(string))
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
		Headers:    DefaultHeaders,
	}
}

// Ok returns 200
func (r *Response) Ok(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusOk, body, options)
}

// Created returns 201
func (r *Response) Created(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusCreated, body, options)
}

// Accepted return 202
func (r *Response) Accepted(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusAccepted, body, options)
}

// NoContent returns 204
func (r *Response) NoContent(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusNoContent, body, options)
}

// BadRequest returns 400
func (r *Response) BadRequest(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusBadRequest, body, options)
}

// Unauthorized returns 401
func (r *Response) Unauthorized(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusUnauthorized, body, options)
}

// Forbidden return 403
func (r *Response) Forbidden(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusForbidden, body, options)
}

// NotFound returns 404
func (r *Response) NotFound(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusNotFound, body, options)
}

// MethodNotAllowed returns 405
func (r *Response) MethodNotAllowed(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusMethodNotAllowed, body, options)
}

// ImATeapot returns 418
func (r *Response) ImATeapot(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusImATeapot, body, options)
}

// UnprocessableEntity returns 422
func (r *Response) UnprocessableEntity(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusUnprocessableEntity, body, options)
}

// TooManyRequests returns 429
func (r *Response) TooManyRequests(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusTooManyRequests, body, options)
}

// ServerError return 500
func (r *Response) ServerError(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusServerError, body, options)
}

// ServiceUnavailable return 503
func (r *Response) ServiceUnavailable(body string, options ...interface{}) events.APIGatewayProxyResponse {
	return r.ProxyResponse(StatusServiceUnavailable, body, options)
}
