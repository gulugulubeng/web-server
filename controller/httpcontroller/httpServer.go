package httpcontroller

import (
	// 第三方路由
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
	"web-server/controller/httpcontroller/usercontoller"
	"web-server/controller/websocketcontroller"
)

// registerHandlers 将一个个小Handler注册进路由器
//func registerHandlers() {
//	// 用户注册Handler
//	http.HandleFunc("/user/register", usercontoller.UserRegister)
//
//	// 用户登录Handler
//	http.HandleFunc("/user/login", usercontoller.UserLogin)
//
//	// 用户信息修改Handler
//	http.HandleFunc("userInfo/update", usercontoller.UserInfoUpdate)
//}

func registerHandlers(router *httprouter.Router) {
	// 用户注册Handler
	router.PUT("/user/register", usercontoller.UserRegister)

	// 用户登录Handler
	router.POST("/user/login", usercontoller.UserLogin)

	// 用户信息修改Handler
	//router.POST("userInfo/update", usercontoller.UserInfoUpdate)

	// 查看文件(http://localhost:8080/static/image/chy.jpg)
	router.ServeFiles("/static/*filepath", http.Dir("./static/"))

	// 将现有的http升级成一个websocket
	router.GET("/ws/:flag/:id", websocketcontroller.UpgradeToWS)
}

// HttpServer 开启http监听并服务
func HttpServer(wg *sync.WaitGroup, httpPort int) {
	// http服务结束
	defer wg.Wait()

	//=================== 开启侦听服务 ====================

	/* 方式一:http.ListenAndServe("localhost:"+strconv.Itoa(httpPort),nil) */
	/* 方式二: 方式二更灵活,更能理解实质发生了什么！方式一底层就是方式二*/
	// http服务器结构体！(对http服务整体控制的对象)

	// 1、路由器
	router := httprouter.New()

	// 2、http服务器
	server := &http.Server{
		Addr:              "localhost:" + strconv.Itoa(httpPort),
		Handler:           router, //路由器
		ReadHeaderTimeout: time.Second * 5,
		IdleTimeout:       time.Minute * 5,
	}
	defer server.Close()

	// 3、给默认路由器注册一个个小Handler
	registerHandlers(router)

	// 4、服务器开始侦听与服务(服务就是使用默认的服务处理器分配http请求！)
	log.Println("开始监听http请求")
	log.Fatalln(server.ListenAndServe())
}
