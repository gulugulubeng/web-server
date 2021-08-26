package service

import (
	"web-server/dao"
	"web-server/domain"
)

func UserRegister(username string, password string) *domain.ResultInfo {
	u := &domain.User{
		Username: username,
		Password: password,
	}
	err := dao.UserRegister(u)
	if err != nil {
		return domain.NewResultInfo(500, "注册失败："+err.Error(), u)
	}
	return domain.NewResultInfo(200, "注册成功！", u)
}
