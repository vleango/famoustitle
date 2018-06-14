package models

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupTest() {
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
