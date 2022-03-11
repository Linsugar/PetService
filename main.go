package main

import (
	"PetService/Middlewares"
	"PetService/MysqlDo"
	"PetService/Routers"
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
	var MapIp = make(map[string]interface{})
	engine()
	Gone.Use(Middlewares.FirstCheck(MapIp), Middlewares.JWThMiddleware())
	Routers.Router(Gone)
	//监听端口默认为8080
	err := Gone.Run(":8000")
	if err != nil {
		return
	}
	defer MysqlDo.Db.Close()
}
