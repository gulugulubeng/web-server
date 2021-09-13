package main

/* 项目全局配置类 */

// 服务配置
const (
	// httpPort http请求端口
	httpPort = 8080
	// tcpPort 聊天请求端口
	tcpPort = 8070
	// udpPort 视频请求端口
	udpPort = 8060
	// multiUdpPort 直播接口
	multiUdpPort = 8050
)

// 数据库配置
const (
	driverName = "mysql" //驱动名必须是相应的数据库

	username = "root"
	password = ""
	protocol = "tcp"
	address  = "127.0.0.1:3306"
	dbname   = "girls"
)

// redis配置
const (
	rAddress   = "172.20.10.4:6379"
	rUsername  = ""
	rPassword  = ""
	rDefaultDB = 0 // use default DB
)

// 日志输出地址
const loggerPath = "C:\\Users\\李想\\Pictures\\Screenshots"
