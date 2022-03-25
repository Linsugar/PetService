package Routers

import (
	"PetService/Views"
	"PetService/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

func Router() {
	R := engine()
	docs.SwaggerInfo.Title = "Swagger专用测试"
	V1Route := R.Group("/v1")
	docs.SwaggerInfo.BasePath = "/v1"
	R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		V1Route.Any("/user", Views.UserController)
		V1Route.Any("/pet", Views.PetController)
		V1Route.Any("/dynamic", Views.DynamicController)
		V1Route.Any("/article", Views.ArticleController)
		V1Route.GET("/weixin", Views.WeixinGet)
		V1Route.POST("/register", Views.Register)
	}
	V2Route := R.Group("/v2")
	{
		V2Route.GET("/user", Views.UserGet)
		V2Route.POST("/user", Views.UserPost)
		V2Route.GET("/pet", Views.PetGet)
		V2Route.POST("/pet", Views.PetPost)
	}

}
