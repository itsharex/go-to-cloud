package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/user"
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

func TestUpdateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	err := UpdateUser(1, &user.User{
		Id:             1,
		Account:        "root",
		RealName:       "肉哦",
		OriginPassword: "123456",
		Email:          "123@email.com",
		Mobile:         "1333333",
	})

	assert.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	err := CreateUser(&user.User{
		Account:        "root",
		RealName:       "肉哦",
		OriginPassword: "123456",
		Email:          "123@email.com",
		Mobile:         "1333333",
	})

	assert.NoError(t, err)
}
