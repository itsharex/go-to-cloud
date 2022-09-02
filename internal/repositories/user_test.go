package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordGenAndCompare(t *testing.T) {
	u := &User{}

	pwd := "OJBK"
	u.SetPassword(&pwd)
	assert.True(t, u.ComparePassword(&pwd))

	pwd1 := "Ojbk"
	assert.True(t, u.ComparePassword(&pwd1))

	pwd2 := "Ojbk "
	assert.False(t, u.ComparePassword(&pwd2))

	pwd3 := ""
	assert.Error(t, u.SetPassword(&pwd3))
	assert.False(t, u.ComparePassword(&pwd3))
}
