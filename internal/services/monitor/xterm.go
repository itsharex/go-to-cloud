package monitor

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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
			err = ws.WriteMessage(mt, message)
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
