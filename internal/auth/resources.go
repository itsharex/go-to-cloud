package auth

import (
	"go-to-cloud/internal/models"
)

var groupPolicies [][]string
var resourcePolicies [][][]string

func init() {
	groupPolicies = [][]string{
		{string(models.Root), "*"},
		{string(models.Ops), string(models.Dev)},
		{string(models.Dev), string(models.Guest)},
	}

	resourcePolicies = [][][]string{
		//{{string(models.Ops), strconv.Itoa(int(models.PodDelete)), "RESOURCE"}},
		//
		//{{string(models.Ops), strconv.Itoa(int(models.PodDelete)), "RESOURCE"}},
		//{{string(models.Ops), strconv.Itoa(int(models.PodShell)), "RESOURCE"}},
		//{{string(models.Ops), strconv.Itoa(int(models.PodViewLog)), "RESOURCE"}},
	}
}
func GroupPolicies() [][]string {
	return groupPolicies
}

func ResourcePolicies() [][][]string {
	return resourcePolicies
}
