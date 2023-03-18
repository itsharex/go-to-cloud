package auth

import (
	"go-to-cloud/internal/models"
	"strconv"
)

func GroupPolicies() [][]string {
	return [][]string{
		{string(models.Root), "*"},
		{string(models.Ops), string(models.Dev)},
		{string(models.Dev), string(models.Guest)},
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
