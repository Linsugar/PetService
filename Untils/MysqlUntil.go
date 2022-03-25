package Untils

import (
	"PetService/Conf"
	"PetService/Models"
	"fmt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", Conf.UserName, Conf.PassWord, Conf.Host, Conf.Port, Conf.Database, Conf.CharSet)
	dataDase, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("有误", err)
		panic(err)

		//return
	}
	Db = dataDase
	var ModelsArrary = []interface{}{&Models.User{}, &Models.PetDetail{}, &Models.Dynamics{}, &Models.Article{}}

	Db.AutoMigrate(ModelsArrary...)
	fmt.Println("链接成功", err)
}
