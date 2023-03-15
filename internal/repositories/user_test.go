package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/user"
	"testing"
	"time"
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

func TestChangePassword(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	account := "test-" + time.Now().Format("20060102150405")
	pwd := "123456"
	u := &user.User{
		Account:        account,
		RealName:       "肉哦",
		OriginPassword: pwd,
		Email:          "123@email.com",
		Mobile:         "1333333",
	}
	err := CreateUser(u)
	assert.NoError(t, err)
	u3 := GetUser(&account, &pwd)

	pwd = "34567"
	err = ResetPassword(u3.ID, &pwd)
	assert.NoError(t, err)
	u2 := GetUser(&account, &pwd)
	assert.True(t, u2.ID == u3.ID)

	pwd2 := "7777"
	err = ResetPasswordWithCheckOldPassword(u3.ID, &pwd, &pwd2)
	u4 := GetUser(&account, &pwd2)
	assert.True(t, u4.ID == u3.ID)
}
