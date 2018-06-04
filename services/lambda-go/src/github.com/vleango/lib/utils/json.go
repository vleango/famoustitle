package utils

import (
	"encoding/json"
)

// JSONStringWithKey String version of text used for body
func JSONStringWithKey(text string, key ...string) string {
	msg := make(map[string]string)
	if len(key) < 1 {
		msg["message"] = text
	} else {
		msg[key[0]] = text
	}

	return MarshalObjectToString(msg)
}

// MarshalObjectToString JSON marshals an interface and returns a json string
func MarshalObjectToString(obj interface{}) string {
	jsonObj, _ := json.Marshal(obj)
	return string(jsonObj)
}
