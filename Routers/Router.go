package Routers

import (
	Views "PetService/ControllerViews"
	"github.com/gin-gonic/gin"
)

func T1(ctx *gin.Context) {
	ctx.Next() //允许下面函数的调用
	//Next-代表的是
	ctx.Abort() //阻止后续的调用
}

func Router(R *gin.Engine) {
	V1Route := R.Group("/v1")
	{
		V1Route.GET("/user", Views.UserController{}.UserGet)
		V1Route.POST("/user", Views.UserController{}.UserPost)
		V1Route.GET("/pet", Views.PetController{}.PetGet)
		V1Route.POST("/pet", Views.PetController{}.PetPost)
		V1Route.POST("/register", Views.RegisterController{}.Register)
		V1Route.GET("/dynamic", Views.DynamicController{}.DynamicAll)
		V1Route.POST("/dynamic", Views.DynamicController{}.DynamicPost)
	}
	V2Route := R.Group("/v2")
	{
		V2Route.GET("/user", Views.UserController{}.UserGet)
		V2Route.POST("/user", Views.UserController{}.UserPost)
		V2Route.GET("/pet", Views.PetController{}.PetGet)
		V2Route.POST("/pet", Views.PetController{}.PetPost)
	}

}
