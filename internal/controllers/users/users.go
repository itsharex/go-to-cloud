package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/conf"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/users"
	"net/http"
	"strconv"
	"strings"
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
			id[i] = user.Id
		}
		response.Success(ctx, id)
	}
}

type tmp struct {
	Users []uint `json:"users"`
	Orgs  []uint `json:"orgs"`
}

// Join
// @Tags User
// @Description 将成员加入/移除组织
// @Success 200
// @Router /api/user/join/{orgId} [put]
// @Param   ContentBody     body     []uint     true  "Request"     example([]uint, userId)
// @Security JWT
func Join(ctx *gin.Context) {
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

	var tmpUser tmp
	if err := ctx.ShouldBindJSON(&tmpUser); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	if err := users.JoinOrg(uint(orgId), tmpUser.Users); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, gin.H{
			"success": true,
		})
	}
}

// ResetPassword
// @Tags User
// @Description 重置用户密码
// @Success 200
// @Router /api/user/{userId}/password/reset [put]
// @Param   ContentBody     body     string     true  "Request"     example(string)
// @Security JWT
func ResetPassword(ctx *gin.Context) {
	exists, _, user, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}
	if !strings.EqualFold(models.RootUserName, *user) {
		msg := "只允许root用户重置密码"
		response.Fail(ctx, http.StatusForbidden, &msg)
		return
	}

	userIdStr := ctx.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	if pwd, err := users.ResetPassword(uint(userId), nil, nil, true); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, gin.H{
			"success":     true,
			"newPassword": *pwd,
		})
	}
}

// AllKinds
// @Tags User
// @Description 所有Kind
// @Accept json
// @Product json
// @Router /api/user/kinds [get]
// @Success 200 {array} string
func AllKinds(ctx *gin.Context) {
	response.Success(ctx, conf.Kinds)
	return
}
