package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	encrypted, err := HashPassword("hogehoge")
	assert.IsType(t, nil, err)
	assert.NotEqual(t, "", encrypted)
	assert.NotEqual(t, nil, encrypted)
}

func TestCheckPasswordHashMatching(t *testing.T) {
	match := CheckPasswordHash("123", "123")
	assert.Equal(t, false, match)
}

func TestCheckPasswordHashNotMatching(t *testing.T) {
	hash, _ := HashPassword("hogehoge")
	match := CheckPasswordHash("hogehoge", hash)
	assert.Equal(t, true, match)
}
