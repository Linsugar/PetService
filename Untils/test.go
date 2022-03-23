package Untils

import (
	"PetService/Conf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func T1(ctx *gin.Context) {
	ctx.Next() //允许下面函数的调用
	//Next-代表的是
	ctx.Abort() //阻止后续的调用
}

type t1 struct {
	gorm.Model
	v1 string `gorm:"unique"`
	v2 int64  `gorm:"not null"`
	v3 map[int]string
}

func main() {
	Conf.Db.AutoMigrate(&t1{})
}
