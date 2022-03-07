package Routers

import (
	"PetService/Views"
	"github.com/gin-gonic/gin"
)


func T1(ctx *gin.Context) {
	ctx.Next() //允许下面函数的调用
	//Next-代表的是
	ctx.Abort() //阻止后续的调用
}

func Router(R *gin.Engine)  {
	V1Route:=R.Group("/v1")
	{
		V1Route.GET("/user",Views.UserController{}.UserGet)
		V1Route.POST("/user",Views.UserController{}.UserPost)
		V1Route.PUT("/user",Views.UserController{}.UserPut)
		V1Route.DELETE("/user",Views.UserController{}.UserDelete)
		V1Route.GET("/pet",Views.PetController{}.PetGet)
		V1Route.POST("/pet",Views.PetController{}.PetPost)
	}
	V2Route:=R.Group("/v2")
	{
		V2Route.GET("/user",Views.UserController{}.UserGet)
		V2Route.POST("/user",Views.UserController{}.UserPost)
		V2Route.GET("/pet",Views.PetController{}.PetGet)
		V2Route.POST("/pet",Views.PetController{}.PetPost)
	}

}
