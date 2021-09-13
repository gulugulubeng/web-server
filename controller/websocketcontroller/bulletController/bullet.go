package bulletController

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Bullet 一个直播间的弹幕结构体
type Bullet struct {
	online int
	m      sync.Map
	lock   sync.RWMutex
}

// AddConn 直播间加入连接
func (b *Bullet) AddConn(conn *websocket.Conn) {
	b.m.Store(conn, nil)
	b.lock.Lock()
	b.online++
	b.lock.Unlock()
}

// DeleteConn 直播间删除、关闭连接
func (b *Bullet) DeleteConn(conn *websocket.Conn) {
	b.m.Delete(conn)
	conn.Close()
	b.lock.Lock()
	b.online--
	b.lock.Unlock()
}

// GetOnline 获取该直播间当前在线人数
func (b *Bullet) GetOnline() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.online
}

// SendMsg 阻塞发送消息
func (b *Bullet) SendMsg(conn *websocket.Conn) error {
	messageType, data, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	b.m.Range(func(key, value interface{}) bool {
		conn = key.(*websocket.Conn) //断言
		conn.WriteMessage(messageType, data)
		return true
	})
	return nil
}
