package Untils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func T1(ctx *gin.Context) {
	ctx.Next() //允许下面函数的调用
	//Next-代表的是
	ctx.Abort() //阻止后续的调用
}

func T2() {
	Testlist := []int{1, 4, 56, 6, 2, 77}
	for i := 0; i < len(Testlist); i++ {
		for j := 0; j < len(Testlist); j++ {
			if Testlist[i] < Testlist[j] {
				Testlist[i], Testlist[j] = Testlist[j], Testlist[i]
			}
		}
	}
	fmt.Println("当前返回的数据：", Testlist)
}
