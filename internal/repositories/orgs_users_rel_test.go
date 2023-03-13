package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/utils"
	"testing"
)

func TestMany2ManyRel(t *testing.T) {

	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	users, err := GetUsersByOrg(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))

	users2, err := GetUsersByOrg(2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users2))
}

func TestMany2ManyRelUpdate(t *testing.T) {

	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	err := UpdateMembersToOrg(1, []uint{3}, []uint{2})
	assert.NoError(t, err)

	users, err := GetUsersByOrg(1)
	o := utils.New[uint]()
	for i := range users {
		utils.Add(o, users[i].ID)
	}
	assert.True(t, utils.Has(o, 3))
	assert.False(t, utils.Has(o, 2))
}
