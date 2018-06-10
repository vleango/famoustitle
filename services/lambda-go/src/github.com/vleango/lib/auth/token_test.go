package auth_test

import (
	"github.com/vleango/lib/auth"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/test"
)

func (suite *TokenSuite) SetupTest() {
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

func (suite *Suite) TestGenerateTokenMissingEmail() {
	token, err := auth.GenerateToken("", "hogehoge")
	suite.Nil(token)
	suite.Equal(auth.ErrMissingParams, err)
}

func (suite *Suite) TestGenerateTokenMissingPass() {
	token, err := auth.GenerateToken("tha.leang@test.com", "")
	suite.Nil(token)
	suite.Equal(auth.ErrMissingParams, err)
}

func (suite *TokenSuite) TestGenerateTokenUserNotFound() {
	token, err := auth.GenerateToken("fake@email.com", "hogehoge")
	suite.Nil(token)
	suite.Equal("record not found", err.Error())
}

func (suite *TokenSuite) TestGenerateTokenPasswordNotMatch() {
	token, err := auth.GenerateToken("tha.leang@test.com", "bad-pass")
	suite.Nil(token)
	suite.Equal("password does not match", err.Error())
}

func (suite *TokenSuite) TestGenerateToken() {
	token, err := auth.GenerateToken("tha.leang@test.com", "hogehoge")
	suite.Nil(err)
	suite.NotNil(token)
}

func (suite *Suite) TestTokenClaimsTokenStringEmpty() {
	claims, err := auth.TokenClaims("")
	suite.Nil(claims)
	suite.Equal("unauthorized", err.Error())
}

func (suite *Suite) TestTokenClaimsTokenStringBadFormat() {
	claims, err := auth.TokenClaims("no-dots")
	suite.Nil(claims)
	suite.Equal("unauthorized", err.Error())
}

func (suite *TokenSuite) TestTokenClaimsWrongSecret() {
	auth.HMACSecret = "different secret"
	claims, err := auth.TokenClaims(suite.userToken)
	suite.Nil(claims)
	suite.Equal("signature is invalid", err.Error())
}

func (suite *TokenSuite) TestTokenClaims() {
	claims, err := auth.TokenClaims(suite.userToken)
	suite.Nil(err)
	suite.Equal("tha.leang@test.com", claims["email"])
	suite.NotNil(claims["exp"])
}
