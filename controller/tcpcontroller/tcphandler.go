package tcpcontroller

import (
	"bufio"
	"log"
	"net"
)

func handleTCPConn(reader *bufio.Reader, conn *net.TCPConn, mes chan string, members map[*net.TCPConn]string) {
	log.Println("连接建立：", conn.RemoteAddr(), "<==>", conn.LocalAddr())

	defer func() {
		delete(members, conn)
		conn.Close()
	}()

	for true {
		str, err := reader.ReadString('\n')
		if err != nil {
			// 读取错误就是连接错误 则直接结束读写！
			log.Println("读取错误", err.Error())
			return
		}
		mes <- members[conn] + ":" + str
	}

}
