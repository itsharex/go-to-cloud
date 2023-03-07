package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type logHeader struct {
	TailLine int64 `json:"tailLine"`
}

func XTermInteractive(ws *websocket.Conn, k8sRepoId uint, containerName string, cancel <-chan struct{}) {

	go func() {
		<-cancel
		fmt.Println("finished")
	}()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read error")
			break
		}
		switch mt {
		case websocket.TextMessage:
			// TODO: 从消息中解析需要获取的日志行号
			var log logHeader
			_ = json.Unmarshal(message, &log)
			err = ws.WriteMessage(mt, []byte("dddddddd"))
		case websocket.PingMessage:
			_ = k8sRepoId
			_ = containerName
			err = ws.WriteMessage(websocket.PongMessage, []byte("pong"))
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
