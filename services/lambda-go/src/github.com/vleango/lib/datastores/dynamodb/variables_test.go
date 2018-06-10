package dynamodb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableNames(t *testing.T) {
	assert.Equal(t, "tech_writer_articles", articleTable)
	assert.Equal(t, "tech_writer_users", userTable)
}
