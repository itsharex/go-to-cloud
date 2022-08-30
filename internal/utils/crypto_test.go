package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAesEny(t *testing.T) {
	plaintext := []byte("Hello中文")

	encoded := AesEny(plaintext)

	decoded := AesEny(encoded)

	assert.Equal(t, plaintext, decoded)
}
