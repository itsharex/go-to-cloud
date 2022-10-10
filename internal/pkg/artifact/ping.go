package artifact

import (
	"fmt"
	"github.com/heroku/docker-registry-client/registry"
	"go-to-cloud/internal/models/artifact"
	"strings"
)

func Ping(testing *artifact.Testing) (bool, error) {
	schema := "https"
	if !testing.IsSecurity {
		schema = "http"
	}

	registryUrl := fmt.Sprintf("%s://%s", schema, strings.TrimSuffix(testing.Url, "/"))
	hub, err := registry.New(registryUrl, testing.User, testing.Password)

	return hub != nil, err
}
