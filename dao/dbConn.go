package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //注册驱动器
)

// 数据库连接操作对象
var db *sql.DB
var err error

// ConnDB 连接数据库
func ConnDB(driverName string, username string, password string, protocol string, address string, dbname string) {

	// 数据源名
	dsn := username + ":" + password + "@" + protocol + "(" + address + ")" + "/" + dbname
	db, err = sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	// 数据库设置
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetConnMaxIdleTime(time.Minute * 2)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// 连接测试
	err = db.Ping()
	if err != nil {
		log.Println("Database:", driverName, "test connected failed！")
	} else {
		fmt.Println("Database:", driverName, "test connected successfully!")
	}
}
