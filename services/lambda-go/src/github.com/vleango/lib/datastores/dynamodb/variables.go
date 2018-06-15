package dynamodb

import (
	"errors"
	"github.com/vleango/config"
)

var (
	articleTable = "famoustitle_articles"
	userTable    = "famoustitle_users"

	svc                     = config.DynamoSvc
	ErrTitleBodyNotProvided = errors.New("missing title and/or body in the HTTP body")
	ErrRecordNotFound       = errors.New("record not found")
)
