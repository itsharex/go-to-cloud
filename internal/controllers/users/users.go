package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/users"
	"net/http"
	"strconv"
)

// Info
// @Tags User
// @Description 查看用户信息
// @Success 200
// @Router /api/user/info [get]
// @Security JWT
func Info(ctx *gin.Context) {
	exists, userId, userName, _, orgs := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	response.Success(ctx, gin.H{
		"userId":   userId,
		"userName": userName,
		"orgs":     orgs,
	})
}

// Logout
// @Tags User
// @Description 注销登录
// @Success 200
// @Router /api/user/logout [get]
// @Security JWT
func Logout(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"code": 20000,
		"data": gin.H{
			"name":   "Hello",
			"avatar": "https://i.jd.com/defaultImgs/9.jpg",
		},
	})
}

// List
// @Tags User
// @Description 列出所有用户
// @Success 200
// @Router /api/user/list [get]
// @Security JWT
func List(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	if u, err := users.GetUserList(); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, u)
	}
}

// Joined
// @Tags User
// @Description 列出加入指定组织的用户
// @Success 200
// @Router /api/user/joined/{orgId} [get]
// @Security JWT
func Joined(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	orgIdStr := ctx.Param("orgId")
	orgId, err := strconv.ParseUint(orgIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	if u, err := users.GetUsersByOrg(uint(orgId)); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		id := make([]uint, len(u))
		for i, user := range u {
			id[i] = user.Key
		}
		response.Success(ctx, id)
	}
}
