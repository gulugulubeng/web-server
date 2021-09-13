package dao

import (
	"context"
	"errors"
	"time"
	"web-server/comment"
	"web-server/model"
)

// UserRegister 用户注册
func UserRegister(u *model.User) error {
	// 查询是否已经注册
	ctx1, _ := context.WithTimeout(comment.Ctx, time.Minute)
	row := db.QueryRowContext(ctx1, "select id from admin where username=?", u.Username)
	err := row.Scan(&u.Id)
	if err == nil {
		return errors.New("此用户已经注册！")
	}
	// 开始注册
	ctx2, _ := context.WithTimeout(comment.Ctx, time.Minute)
	result, err := db.ExecContext(ctx2, "insert into admin(username,password) values(?,?)", u.Username, u.Password)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	u.Id = int(id)
	return nil
}
