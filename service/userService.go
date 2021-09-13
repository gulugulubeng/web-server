package service

import (
	"web-server/dao"
	"web-server/model"
)

func UserRegister(username string, password string) *model.ResultInfo {
	u := &model.User{
		Username: username,
		Password: password,
	}
	err := dao.UserRegister(u)
	if err != nil {
		return model.NewResultInfo(500, "注册失败："+err.Error(), u)
	}
	return model.NewResultInfo(200, "注册成功！", u)
}
