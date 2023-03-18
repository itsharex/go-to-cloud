package auth

import (
	"go-to-cloud/conf"
	"go-to-cloud/internal/models"
	"strconv"
)

func GroupPolicies() [][]string {
	return [][]string{
		{string(conf.Root), "*"},
	}
}

func ResourcePolicies() [][][]string {
	return [][][]string{
		{{string(conf.Ops), strconv.Itoa(int(models.PodDelete)), "RESOURCE"}},
		{{string(conf.Ops), strconv.Itoa(int(models.PodShell)), "RESOURCE"}},
		{{string(conf.Ops), strconv.Itoa(int(models.PodViewLog)), "RESOURCE"}},
	}
}
