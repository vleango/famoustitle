package utils_test

import (
  "testing"
  "github.com/stretchr/testify/suite"
  "github.com/vleango/lib/utils"
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
  uniqueArray := utils.RemoveStringDuplicatesUnordered([]string{"ruby", "rails", "ruby"})
  suite.Equal(2, len(uniqueArray))
  suite.Contains(uniqueArray, "ruby")
  suite.Contains(uniqueArray, "rails")
}
