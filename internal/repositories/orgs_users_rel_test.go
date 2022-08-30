package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMany2ManyRel(t *testing.T) {

	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	users, err := FetchUsersByOrg(2)
	assert.NoError(t, err)
	assert.NotNil(t, users)
}
