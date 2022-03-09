package main

import (
	"PetService/Middlewares"
	"PetService/MysqlDo"
	"PetService/Routers"
	"github.com/gin-gonic/gin"
)

func main() {
	var MapIp = make(map[string]interface{})
	r := gin.Default()
	r.Use(Middlewares.FirstCheck(MapIp), Middlewares.JWThMiddleware())
	Routers.Router(r)
	//监听端口默认为8080
	err := r.Run(":8000")
	if err != nil {
		return
	}
	defer MysqlDo.Db.Close()
}
