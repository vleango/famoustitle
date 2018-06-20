package auth_test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vleango/lib/auth"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
	"time"
)

func (suite *AuthenticateSuite) SetupTest() {
	test.CleanDataStores()
	tokens := test.CreateUserTable(map[string]interface{}{
		"user": models.User{
			FirstName: "Tha",
			LastName:  "Leang",
			Email:     "tha.leang@test.com",
		},
		"password": "hogehoge",
	})
	suite.userToken = tokens[0]
}

func (suite *Suite) TestAuthenticateUserEmptyToken() {
	user, err := auth.AuthenticateUser("Bearer")
	suite.Nil(user)
	suite.Equal(auth.ErrAuthTokenUnauthorized, err)
}

func (suite *Suite) TestAuthenticateUserNonJWTFormat() {
	user, err := auth.AuthenticateUser("Bearer bad-format")
	suite.Nil(user)
	suite.Equal(auth.ErrAuthTokenUnauthorized, err)
}

func (suite *Suite) TestAuthenticateUserClaimNotValid() {
	user, err := auth.AuthenticateUser("Bearer this.not.valid")
	suite.Nil(user)
	suite.Equal(auth.ErrAuthTokenUnauthorized, err)
}

func (suite *Suite) TestAuthenticateUserClaimExpired() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": "fake@email.com",
		"exp":   time.Now().Unix() - 10000,
	})

	tokenString, _ := token.SignedString([]byte(auth.HMACSecret))
	user, err := auth.AuthenticateUser(fmt.Sprintf("Bearer %v", tokenString))
	suite.Nil(user)
	suite.Equal(auth.ErrAuthTokenExpired, err)
}

func (suite *Suite) TestAuthenticateUserEmptyBearer() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": "fake@email.com",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(auth.HMACSecret))
	user, err := auth.AuthenticateUser(tokenString)
	suite.Nil(user)
	suite.Equal(auth.ErrAuthTokenUnauthorized, err)
}

func (suite *AuthenticateSuite) TestAuthenticateUserEmailNotFound() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": "fake@email.com",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(auth.HMACSecret))
	user, err := auth.AuthenticateUser(fmt.Sprintf("Bearer %v", tokenString))
	suite.Nil(user)
	suite.Equal("record not found", err.Error())
}

func (suite *AuthenticateSuite) TestAuthenticateUser() {
	user, err := auth.AuthenticateUser(fmt.Sprintf("Bearer %v", suite.userToken))
	suite.Equal("Tha", user.FirstName)
	suite.Equal("Leang", user.LastName)
	suite.Equal("tha.leang@test.com", user.Email)
	suite.NotNil(user.ID)
	suite.Nil(err)
}
