package main

import (
	"PetService/Middlewares"
	"PetService/Routers"
	"PetService/Tasks"
	"PetService/Untils"
)

//var MapIp = map[string]interface{}{}

func main() {
	Routers.Router()
	Routers.Gone.Use(Middlewares.JWThMiddleware())
	////Gone.Use(Middlewares.FirstCheck(MapIp), Middlewares.JWThMiddleware())
	//Gone.Use(Middlewares.JWThMiddleware())
	//监听端口默认为8080

	err := Routers.Gone.Run(":8000")
	if err != nil {
		return
	}
	Tasks.TaskInitAll()
	defer Untils.Db.Close()
	//每天凌晨1点执行一次

}
