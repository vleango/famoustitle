package logs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableNames(t *testing.T) {
	debug := DebugMessage(100, "hello")
	assert.Equal(t, "********** DEBUG: (100) hello", debug)
}
