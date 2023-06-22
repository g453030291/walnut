package rds

import "github.com/redis/go-redis/v9"

var Rds *redis.Client

func Init() {
	Rds = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	print("redis init success\n")
}
