package websocketcontroller

import (
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
	"web-server/controller/websocketcontroller/bulletController"
)

var (
	upgrader = &websocket.Upgrader{
		HandshakeTimeout: time.Second * 5,
		ReadBufferSize:   2048,
		WriteBufferSize:  2048,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// UpgradeToWS 将HTTP升级为WebSocket协议
func UpgradeToWS(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println("连接错误")
		return
	}

	switch params.ByName("flag") {
	case "bullet":
		bulletController.BulletCtr(conn, params)
	}

}
