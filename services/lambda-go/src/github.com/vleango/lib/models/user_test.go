package models

import (
	"time"
)

func (suite *Suite) TestUser() {
	user := User{}
	suite.IsType("testing-2134", user.ID)
	suite.IsType("Tha", user.FirstName)
	suite.IsType("Leang", user.LastName)
	suite.IsType("email", user.Email)
	suite.IsType("digest", user.PasswordDigest)
	suite.IsType(true, user.IsAdmin)
	suite.IsType(true, user.IsWriter)
	suite.IsType(map[string]string{}, user.Articles)
	suite.IsType(time.Now(), user.CreatedAt)
	suite.IsType(time.Now(), user.UpdatedAt)
}

func (suite *Suite) TestFullName() {
	user := User{
		FirstName: "Tha",
		LastName:  "Leang",
	}

	suite.Equal("Tha Leang", user.FullName())
}
