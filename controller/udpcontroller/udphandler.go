package udpcontroller

import (
	"log"
	"net"
	"strings"
)

func handleUDPConn(rAddr *net.UDPAddr, data []byte) {
	str := string(data)
	switch {
	case strings.HasPrefix(str, "file:"):
		udpFileDownload(rAddr, strings.TrimLeft(str, "file:"))
	case strings.HasPrefix(str, "mv4:"):
		udpMv4Play(rAddr, strings.TrimLeft(str, "mv4:"))
	default:
		log.Println(str, "udp请求错误！")
	}
}

func udpFileDownload(raddr *net.UDPAddr, filePath string) {

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Println("udp连接错误", err.Error())
		return
	}
	defer conn.Close()

}

func udpMv4Play(raddr *net.UDPAddr, mv4Name string) {
}
