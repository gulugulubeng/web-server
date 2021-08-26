package httpcontroller

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
	"web-server/controller/httpcontroller/usercontoller"
)

// registerHandlers 将一个个小Handler注册进路由器
func registerHandlers() {
	// 用户注册Handler
	http.HandleFunc("/user/register", usercontoller.UserRegister)

	// 用户登录Handler
	http.HandleFunc("/user/login", usercontoller.UserLogin)

	// 用户信息修改Handler
	http.HandleFunc("userInfo/update", usercontoller.UserInfoUpdate)
}

// HttpServer 开启http监听并服务
func HttpServer(wg *sync.WaitGroup, httpPort int) {
	// http服务结束
	defer wg.Wait()

	//=================== 开启侦听服务 ====================

	/* 方式一:http.ListenAndServe("localhost:"+strconv.Itoa(httpPort),nil) */
	/* 方式二: 方式二更灵活,更能理解实质发生了什么！方式一底层就是方式二*/
	// http服务器结构体！(对http服务整体控制的对象)
	server := http.Server{
		Addr:              "localhost:" + strconv.Itoa(httpPort),
		Handler:           nil, //nil采用默认服务处理器(路由器)
		ReadHeaderTimeout: time.Minute,
	}
	defer server.Close()

	// 给默认路由器注册一个个小Handler
	registerHandlers()

	// 服务器开始侦听与服务(服务就是使用默认的服务处理器分配http请求！)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("http服务开启失败:", err.Error())
	}

}
