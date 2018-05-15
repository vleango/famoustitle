package utils

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

func (suite *Suite) TestRemoveStringDuplicatesUnordered() {
	uniqueArray := RemoveStringDuplicatesUnordered([]string{"ruby", "rails", "ruby"})
	suite.Equal(2, len(uniqueArray))
	suite.Contains(uniqueArray, "ruby")
	suite.Contains(uniqueArray, "rails")
}
