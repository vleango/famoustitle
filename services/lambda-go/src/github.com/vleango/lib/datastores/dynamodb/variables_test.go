package dynamodb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableNames(t *testing.T) {
	assert.Equal(t, "famoustitle_articles", articleTable)
	assert.Equal(t, "famoustitle_users", userTable)
}
