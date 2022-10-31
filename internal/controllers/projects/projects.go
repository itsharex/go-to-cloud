package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	project2 "go-to-cloud/internal/models/project"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
)

// List
// @Tags Projects
// @Description 查看项目信息
// @Success 200 {array} project.DataModel
// @Router /api/projects/list [get]
// @Security JWT
func List(ctx *gin.Context) {
	exists, _, _, orgs, _ := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	m, err := project.List(orgs)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
	}
}

// Create
// @Tags Projects
// @Description 创建新的项目
// @Param   ContentBody     body     project.DataModel     true  "Request"     example(project.DataModel)
// @Success 200
// @Router /api/projects/create [POST]
// @Security JWT
func Create(ctx *gin.Context) {
	var req project2.DataModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	exists, userId, _, orgs, _ := util.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	_, err := project.CreateNewProject(userId, orgs, req)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
