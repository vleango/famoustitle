package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/password"
	"strings"
	"time"
)

var (
	// TODO
	HMACSecret = "this is a test"

	ErrMissingParams = fmt.Errorf("missing required params")
)

func GenerateToken(email string, pass string) (*models.User, *string, error) {
	if email == "" || pass == "" {
		return nil, nil, ErrMissingParams
	}

	// find user by email
	user, err := dynamodb.UserFindByEmail(email)
	if err != nil {
		return nil, nil, err
	}

	// get the user's password_digest // check if password match
	ok := password.CheckPasswordHash(pass, user.PasswordDigest)
	if !ok {
		return nil, nil, fmt.Errorf("password does not match")
	}

	// generate jwt
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(HMACSecret))

	// return the jwt
	return user, &tokenString, nil
}

func TokenClaims(tokenString string) (map[string]interface{}, error) {
	if tokenString == "" || strings.Count(tokenString, ".") != 2 {
		return nil, fmt.Errorf("unauthorized")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(HMACSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
