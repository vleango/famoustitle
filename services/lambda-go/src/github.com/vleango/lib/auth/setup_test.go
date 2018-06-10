package auth_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

type AuthenticateSuite struct {
	suite.Suite
	userToken string
}

type TokenSuite struct {
	suite.Suite
	userToken string
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
	suite.Run(t, new(AuthenticateSuite))
	suite.Run(t, new(TokenSuite))
}
