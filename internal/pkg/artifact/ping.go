package artifact

import (
	"fmt"
	"go-to-cloud/internal/models/artifact"
	"net/http"
	"strings"
)

func Ping(testing *artifact.Testing) (bool, error) {
	schema := "https"
	if !testing.IsSecurity {
		schema = "http"
	}

	registryUrl := fmt.Sprintf("%s://%s", schema, strings.TrimSuffix(testing.Url, "/"))
	_, err := http.Get(registryUrl)
	if err != nil {
		return false, err
	}

	// TODO: check if 401
	return false, nil
}
