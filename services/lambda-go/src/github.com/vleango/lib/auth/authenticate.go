package auth

import (
	"fmt"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/logs"
	"github.com/vleango/lib/models"
	"strings"
)

var (
	ErrAuthTokenUnauthorized = fmt.Errorf("unauthorized")
	ErrAuthTokenExpired      = fmt.Errorf("token is expired")
)

func AuthenticateUser(bearer string) (*models.User, error) {
	if !strings.HasPrefix(bearer, "Bearer") {
		return nil, ErrAuthTokenUnauthorized
	}

	token := strings.TrimPrefix(bearer, "Bearer ")
	claims, err := TokenClaims(token)
	if err != nil {
		logs.DebugMessage(400, err.Error())
		if strings.ToLower(err.Error()) == ErrAuthTokenExpired.Error() {
			return nil, ErrAuthTokenExpired
		} else {
			return nil, ErrAuthTokenUnauthorized
		}
	}

	// get user
	return dynamodb.UserFindByEmail(claims["email"].(string))
}
