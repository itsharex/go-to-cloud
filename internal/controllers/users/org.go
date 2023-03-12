package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/users"
	"net/http"
)

// OrgList
// @Tags User
// @Description 查看组织
// @Success 200
// @Router /api/user/org/list [get]
// @Security JWT
func OrgList(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	if orgs, err := users.GetOrgList(); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, orgs)
	}
}
