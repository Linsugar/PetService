package main

import (
	"PetService/MysqlDo"
	"PetService/Routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Routers.Router(r)
	//监听端口默认为8080
	err := r.Run(":8000")
	if err != nil {
		return
	}
	defer MysqlDo.Db.Close()
}
