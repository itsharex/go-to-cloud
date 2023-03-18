package auth

import (
	"go-to-cloud/internal/models"
	"strconv"
)

func GroupPolicies() [][]string {
	return [][]string{
		{string(models.Root), "*"},
	}
}

func ResourcePolicies() [][][]string {
	return [][][]string{
		{{string(models.Ops), strconv.Itoa(int(models.PodDelete)), "RESOURCE"}},

		{{string(models.Ops), strconv.Itoa(int(models.PodDelete)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.PodShell)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.PodViewLog)), "RESOURCE"}},
	}
}
