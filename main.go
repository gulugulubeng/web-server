package main

import (
	"log"
	"sync"
	"web-server/controller/httpcontroller"
	"web-server/controller/tcpcontroller"
	"web-server/dao"
)

// 程序入口
func main() {
	log.Println("=======================项目开启========================")
	// main函数的等待组
	wg := new(sync.WaitGroup)

	// 开启http监听并服务
	wg.Add(1)
	go httpcontroller.HttpServer(wg, httpPort)
	// 开启tcp监听并服务
	wg.Add(1)
	go tcpcontroller.TCPServer(wg, tcpPort)
	// 开启udp监听并服务
	//wg.Add(1)
	//go udpcontroller.UDPServer(wg,udpPort)

	// 连接数据库
	dao.ConnDB(driverName, username, password, protocol, address, dbname)
	// 连接Redis
	//dao.ConnRDB(rAddress,rUsername,rPassword, rDefaultDB)

	wg.Wait()
	log.Println("=======================项目注销========================")
}
