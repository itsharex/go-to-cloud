package registry

import (
	"fmt"
	"github.com/heroku/docker-registry-client/registry"
	"strings"
)

func GetRegistryHub(isSecurity bool, url, user, password *string) (hub *registry.Registry, err error) {
	schema := "https"
	if !isSecurity {
		schema = "http"
	}

	registryUrl := fmt.Sprintf("%s://%s", schema, strings.TrimSuffix(*url, "/"))
	hub, err = registry.New(registryUrl, *user, *password)

	return
}
