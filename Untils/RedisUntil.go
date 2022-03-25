package Untils

import (
	"PetService/Conf"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Conn *redis.Client
var Ctx = context.Background()

func init() {
	opt := redis.Options{
		Addr: Conf.Addr,
	}
	Conn = redis.NewClient(&opt)
}

func GetRedisValue(key string) interface{} {
	//根据key获取值
	res := Conn.Get(Ctx, key)
	value, err := res.Result()
	if err != nil {
		fmt.Println("出现错误", err)
		if err == redis.Nil {
			fmt.Println("没有值")
			return nil
		}
	} else {
		fmt.Println("\nnim\n", value)
	}
	return value
}

func SetRedisValue(key string, value interface{}, t time.Duration) {
	//根据key设置值
	res := Conn.Set(Ctx, key, value, t)
	fmt.Println("当前Set返回值：res===", res)

}
