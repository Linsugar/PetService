package main

import (
	"PetService/Conf"
	"PetService/Middlewares"
	"PetService/Routers"
	"PetService/Untils"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once
var Gone *gin.Engine

//实现单例只创建一次
func engine() *gin.Engine {
	once.Do(func() {
		Gone = gin.Default()
	})
	return Gone
}

func main() {
	engine()
	//Gone.Use(Middlewares.FirstCheck(MapIp), Middlewares.JWThMiddleware())
	Gone.Use(Middlewares.JWThMiddleware())
	Routers.Router(Gone)
	//监听端口默认为8080
	err := Gone.Run(":8000")

	if err != nil {
		return
	}
	defer Conf.Db.Close()
	Untils.Cr.Start()
	//每天凌晨1点执行一次

}
