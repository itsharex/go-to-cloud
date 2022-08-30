package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchInfrastructures(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	infra, err := FetchInfrastructures(1, InfraTypeK8s)
	assert.NoError(t, err)
	assert.Greater(t, len(infra), 0)
	for _, inf := range infra {
		assert.Equal(t, InfraTypeK8s, inf.Type)
	}
}
