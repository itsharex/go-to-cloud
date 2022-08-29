package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchInfrastructures(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to ci is seperated from DB")
	}

	infra := FetchInfrastructures(1, K8s)

	assert.NotEmpty(t, infra)
}
