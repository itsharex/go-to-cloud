package monitor

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/monitor"
	"net/http"
)

// DisplayLog 进入容器内部执行命令行交互
func DisplayLog(ctx *gin.Context) {

	k8sRepoId, err := getUIntParamFromQueryOrPath("k8s", ctx, false)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	container := ctx.Param("container") // 允许为空，空时进入默认第一个容器内部

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	ws, err := monitor.Upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	defer func() {
		ws.Close()
	}()

	monitor.XTermInteractive(ws, k8sRepoId, container, ctx.Done())
}
