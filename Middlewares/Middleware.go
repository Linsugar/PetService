package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func FirstCheck(MapValue map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		IP := c.ClientIP()
		v, ok := MapValue[IP]
		if ok {
			t1 := v.(int64)
			if t2 := time.Now().Unix(); t2-t1 > 5 {
				//可以继续操作
				c.Next()
			} else {
				//
				c.Abort()
			}
		} else {
			MapValue[IP] = time.Now().Unix()
			fmt.Printf("当前ip名单：%v", MapValue)
			c.Next()
		}
	}
}

func JWThMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的token中
		url := c.Request.URL
		fmt.Printf("当前路由：%v", url)
		method := c.Request.Method
		fmt.Printf("当前方法：%v", method)
		if url.Path == "/v1/user" && method == "POST" {
			c.Next()
			return
		}
		token := c.Request.Header.Get("token")
		if token == "" {
			// 处理 没有token的时候
			c.JSON(403, gin.H{
				"err": "丢失token",
			})
			c.Abort() // 不会继续停止
			return
		}
		// 解析
		mc, err := ParseToken(token)
		if err != nil {
			// 处理 解析失败
			c.Abort()
			return
		}
		fmt.Println("解析成功")
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userID", mc.UserID)
		c.Next()
	}
}
