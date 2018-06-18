package dynamodb

import (
	"errors"
	"github.com/vleango/config"
)

var (
	articleTable = config.DynamoArticlesTable
	userTable    = config.DynamoUsersTable

	svc                     = config.DynamoSvc
	ErrTitleBodyNotProvided = errors.New("missing title and/or body in the HTTP body")
	ErrRecordNotFound       = errors.New("record not found")
)
