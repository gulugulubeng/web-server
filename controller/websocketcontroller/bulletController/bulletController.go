package bulletController

import (
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"log"
	"strconv"
	"time"
)

var bullets = NewBullets()

// BulletCtr 直播间弹幕控制
func BulletCtr(conn *websocket.Conn, params httprouter.Params) {
	// 解析直播间ID
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Println("id解析成错误")
		return
	}
	// 根据ID获取直播间
	b := bullets.GetBullet(id)
	if b == nil {
		b = bullets.AddBullet(id)
	}

	// 直播间添加连接
	b.AddConn(conn)

	// 发送心跳
	go func() {
		for true {
			data := "heartbeat\n当前直播间ID:" + strconv.Itoa(id) + "\n当前直播间在线人数：" + strconv.Itoa(b.GetOnline())
			if err = conn.WriteMessage(1, []byte(data)); err != nil {
				return
			}
			time.Sleep(4 * time.Second)
		}
	}()

	// 广播消息
	for true {
		if err = b.SendMsg(conn); err != nil {
			break
		}
	}

	// 直播间移除连接
	b.DeleteConn(conn)
}
