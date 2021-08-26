package dao

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"web-server/comment"
)

// redis连接客户端
var rdb *redis.Client

// ConnRDB 连接redis
func ConnRDB(address string, username string, password string, defaultDB int) {
	// redis客户端
	rdb = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     address,
		Username: username,
		Password: password,  // no password set
		DB:       defaultDB, // use default DB
	})

	// 测试连接
	status := rdb.Ping(comment.Ctx)
	if status.Err() != nil {
		log.Println("redis test connected failed！")
	} else {
		fmt.Println("redis test connected successfully！")
	}

	err = rdb.Set(comment.Ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(comment.Ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(comment.Ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
