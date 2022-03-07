package Untils

import "github.com/gin-gonic/gin"

func T1(ctx *gin.Context) {
	ctx.Next() //允许下面函数的调用
	//Next-代表的是
	ctx.Abort() //阻止后续的调用
}


