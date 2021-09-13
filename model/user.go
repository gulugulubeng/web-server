package model

// User 用户结构体
type User struct {
	Id       int    `json:"id,1"`
	Username string `json:"username,2"`
	Password string `json:"password,3"`
}
