package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordGenAndCompare(t *testing.T) {
	u := &User{}

	pwd := "OJBK"
	u.SetPassword(&pwd)
	assert.True(t, u.comparePassword(&pwd))

	pwd = "OJBK"
	u.SetPassword(&pwd)
	assert.True(t, u.comparePassword(&pwd))

	pwd1 := "Ojbk"
	assert.True(t, u.comparePassword(&pwd1))

	pwd2 := "Ojbk "
	assert.False(t, u.comparePassword(&pwd2))

	pwd3 := ""
	assert.Error(t, u.SetPassword(&pwd3))
	assert.False(t, u.comparePassword(&pwd3))
}
