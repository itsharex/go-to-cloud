package monitor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

var upgrade = websocket.Upgrader{}

// Cli 进入容器内部执行命令行交互
func Cli(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	k8sRepoId, err := getUIntParamFromQueryOrPath("k8s", ctx, false)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	container := ctx.Param("container") // 允许为空，空时进入默认第一个容器内部

	ws, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}
	defer ws.Close()

	go func() {
		<-ctx.Done()
	}()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read error")
			break
		}
		switch mt {
		case websocket.TextMessage:
			err = ws.WriteMessage(mt, message)
		case websocket.PingMessage:
			_ = k8sRepoId
			_ = container
			err = ws.WriteMessage(websocket.PongMessage, []byte("pong"))
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
