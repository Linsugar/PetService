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
	V1Route := R.Group("/UserCenter")
	docs.SwaggerInfo.BasePath = "/UserCenter"
	R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		V1Route.Any("/user", Views.UserController)
		V1Route.Any("/pet", Views.PetController)
		V1Route.Any("/dynamic", Views.DynamicController)
		V1Route.Any("/article", Views.ArticleController)
		V1Route.GET("/weixin", Views.WeixinGet)
		V1Route.POST("/register", Views.Register)
	}
	V2Route := R.Group("/UserConfig")
	{

		V2Route.POST("/QiNiu", Views.SetQINiuToken) //获取七牛云token
		V2Route.POST("/Code", Views.CodeController) //获取验证码
	}

}
