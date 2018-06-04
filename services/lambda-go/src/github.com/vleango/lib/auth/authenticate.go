package auth

import (
	"fmt"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/models"
	"strings"
	"time"
)

func AuthenticateUser(bearer string) (*models.User, error) {
	token := strings.TrimPrefix(bearer, "Bearer ")

	claims, err := TokenClaims(token)
	if err != nil {
		return nil, err
	}

	// check if expired
	if claims["exp"] == nil || time.Now().Unix() > int64(claims["exp"].(float64)) {
		return nil, fmt.Errorf("expired token")
	}

	// get user
	return dynamodb.UserFindByEmail(claims["email"].(string))
}
