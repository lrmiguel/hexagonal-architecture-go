package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	assert.Equal(t, `{"message":"Hello Json"}`, string(result))
}
