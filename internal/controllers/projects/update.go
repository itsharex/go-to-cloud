package projects

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	project2 "go-to-cloud/internal/models/project"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// UpdateProject 更新项目信息
// @Tags Projects
// @Description 更新项目信息
// @Success 200
// @Param   ContentBody     body     project.DataModel     true  "Request"     example(project.DataModel)
// @Router /api/projects [PUT]
// @Security JWT
func UpdateProject(ctx *gin.Context) {
	var req project2.DataModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if req.OrgId < 0 {
		response.BadRequest(ctx, errors.New("one organization at least"))
		return
	}

	exists, userId, _, orgs, _ := utils.CurrentUser(ctx)
	_ = userId
	_ = orgs
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	// TODO: 更新（归属组织不可更改)
	//_, err := project.CreateNewProject(userId, orgs, req)
	//if err != nil {
	//	msg := err.Error()
	//	response.Fail(ctx, http.StatusInternalServerError, &msg)
	//	return
	//}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
