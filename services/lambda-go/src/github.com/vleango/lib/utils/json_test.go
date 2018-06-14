package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONStringWithKeyPresent(t *testing.T) {
	msg := JSONStringWithKey("hello", "bye")
	assert.Equal(t, "{\"bye\":\"hello\"}", msg)
}

func TestJSONStringWithKeyNotPresent(t *testing.T) {
	msg := JSONStringWithKey("hello")
	assert.Equal(t, "{\"message\":\"hello\"}", msg)
}

func TestMarshalObjectToString(t *testing.T) {
	obj := map[string]string{
		"hello": "baby",
		"bye":   "hi",
	}

	msg := MarshalObjectToString(obj)
	assert.Equal(t, "{\"bye\":\"hi\",\"hello\":\"baby\"}", msg)
}
