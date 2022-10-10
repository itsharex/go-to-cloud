package artifact

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import artifactModel "go-to-cloud/internal/models/artifact"

func TestPing_Should_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ok, err := Ping(&artifactModel.Testing{
		IsSecurity: false,
		Url:        "81.68.216.88:8080",
		User:       "admin",
		Password:   "some password",
	})

	assert.True(t, ok)
	assert.NoError(t, err)
}
