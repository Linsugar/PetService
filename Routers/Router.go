package Routers

import (
	Views "PetService/Views"
	docs "PetService/docs"
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
		V1Route.GET("/user", Views.UserGet)
		V1Route.POST("/user", Views.UserPost)
		V1Route.GET("/pet", Views.PetGet)
		V1Route.POST("/pet", Views.PetPost)
		V1Route.POST("/register", Views.Register)
		V1Route.GET("/dynamic", Views.DynamicAll)
		V1Route.POST("/dynamic", Views.DynamicPost)
		V1Route.GET("/article", Views.ArticleAll)
		V1Route.GET("/weixin", Views.WeixinGet)
		V1Route.POST("/article", Views.ArticlePost)
	}
	V2Route := R.Group("/v2")
	{
		V2Route.GET("/user", Views.UserGet)
		V2Route.POST("/user", Views.UserPost)
		V2Route.GET("/pet", Views.PetGet)
		V2Route.POST("/pet", Views.PetPost)
	}

}
