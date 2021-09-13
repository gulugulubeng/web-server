package udpcontroller

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

func UDPServer(wg *sync.WaitGroup, udpPort int) {
	defer wg.Done()

	// 解析成地址对象
	lAddr, _ := net.ResolveUDPAddr("udp", "localhost:"+strconv.Itoa(udpPort))

	// 创建本地udp读取套接地址
	conn, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Println("udp服务监听器创建失败:", err.Error())
		return
	}
	defer conn.Close()

	// 读取缓存中的数据
	buf := make([]byte, 0, 1024*2)
	var rAddr *net.UDPAddr
	for {
		_, rAddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Println("udp数据读取错误：", err.Error())
			return
		}
		if rAddr != nil {
			break
		}
	}
	// 获取要下载的文件
	file, err := os.Open(string(buf))
	if err != nil {
		log.Println("请求文件错误：", err.Error())
	}
	sendFile(file, rAddr, conn)
}

func sendFile(f *os.File, rAddr *net.UDPAddr, conn *net.UDPConn) {
	defer f.Close()

	// 将文件读取到用户空间的缓冲区
	buf := make([]byte, 0, 1024*2)

	for true {
		num, err := f.Read(buf[0:])
		if err != nil {
			log.Println("文件读取错误：", err.Error())
			return
		}
		_, err = conn.WriteToUDP(buf, rAddr)
		if err != nil {
			return
		}
		if num == 0 {
			fmt.Println("发送结束")
			return
		}
	}
}
