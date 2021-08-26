package tcpcontroller

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

// TCPServer 开启tcp监听并服务
func TCPServer(wg *sync.WaitGroup, tcpPort int) {
	defer wg.Done()

	// 解析成地址对象
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:"+strconv.Itoa(tcpPort))
	// 创建tcp监听器
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println("tcp服务监听器创建失败:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("开始监听tcp请求")

	// 消息队列
	mes := make(chan string, 5)

	// 在线成员
	members := make(map[*net.TCPConn]string)

	// 广播信息
	go broadCastMes(members, mes)

	var reader *bufio.Reader

	// 开始循环监听
	for true {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("tcp连接异常：", err.Error())
			continue
		}
		// 1.有一个连接
		reader = bufio.NewReader(conn)
		// 2.读取连接用户信息
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)
		members[conn] = id
		id = id + ":" + id + "进入聊天室！\n"
		mes <- id
		// 单独开启对话
		go handleTCPConn(reader, conn, mes, members)
	}
}

func broadCastMes(members map[*net.TCPConn]string, mes chan string) {
	for true {
		select {
		case str := <-mes:
			sender := strings.Split(str, ":")[0]
			message := []byte(str)
			for conn, id := range members {
				if id != sender {
					conn.Write(message)
				}
			}
		}
	}
}
