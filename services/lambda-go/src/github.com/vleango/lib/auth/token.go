package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/password"
	"time"
)

var (
	// TODO
	hmacSecret = "this is a test"
)

func GenerateToken(email string, pass string) (*string, error) {
	// find user by email
	user, err := dynamodb.UserFindByEmail(email)
	if err != nil {
		return nil, err
	}

	// get the user's password_digest // check if password match
	ok := password.CheckPasswordHash(pass, user.PasswordDigest)
	if !ok {
		return nil, fmt.Errorf("password does not match")
	}

	// generate jwt
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSecret))

	// return the jwt
	return &tokenString, nil
}

func TokenClaims(tokenString string) (map[string]interface{}, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("unauthorized")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hmacSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
