package usercontoller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"web-server/model"
	"web-server/service"
)

// UserRegister 注册函数
func UserRegister(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	username := req.FormValue("username")
	password := req.FormValue("password")

	// 统一类型返回结果
	result := new(model.ResultInfo)

	// 参数校验
	if username != "" && password != "" {
		result = service.UserRegister(username, password)
	} else {
		result = model.NewResultInfo(403, "注册失败！用户名或密码错误！", nil)
	}

	// 将结果实例序列化成json格式
	ans, err := json.Marshal(result)
	if err != nil {
		log.Println("统一数据类型序列化错误：", err)
	}
	res.Write(ans)
}

// UserLogin 登录函数
func UserLogin(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {}
