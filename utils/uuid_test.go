package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UUID(t *testing.T) {
	uuid := UUID()
	assert.NotEmpty(t, uuid)
}

func Test_UUIDWithPrefix(t *testing.T) {
	uuid := UUIDWithPrefix("p")
	assert.NotEmpty(t, uuid)
}
