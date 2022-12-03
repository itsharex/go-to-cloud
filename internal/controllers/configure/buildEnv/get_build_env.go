package buildEnv

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/pkg/response"
)

type envGroup struct {
	Label   string `json:"label"`
	Options []struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"options"`
}

// BuildEnv 构建环境
// @Tags BuildConfigure
// @Description 构建环境
// @Success 200 {array} envGroup
// @Router /api/configure/build/env [get]
// @Security JWT
func BuildEnv(ctx *gin.Context) {
	response.Success(ctx, []envGroup{
		{
			Label: ".Net",
			Options: []struct {
				Value string `json:"value"`
				Label string `json:"label"`
			}{
				{
					Value: dotNet3,
					Label: ".NET Core 3.1",
				}, {
					Value: dotNet5,
					Label: ".NET 5",
				}, {
					Value: dotNet6,
					Label: ".NET 6",
				}, {
					Value: dotNet7,
					Label: ".NET 7",
				},
			},
		},
		{
			Label: "Golang",
			Options: []struct {
				Value string `json:"value"`
				Label string `json:"label"`
			}{
				{
					Value: go116,
					Label: "Go 1.16",
				}, {
					Value: go117,
					Label: "Go 1.17",
				}, {
					Value: go118,
					Label: "Go 1.18",
				}, {
					Value: go119,
					Label: "Go 1.19",
				},
			},
		},
	})
}
