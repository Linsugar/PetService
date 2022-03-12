package MysqlDo

import (
	"PetService/Models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	host := "139.155.88.241"
	port := "3388"
	database := "oneweb"
	username := "root"
	password := "123456"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	dataDase, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("有误", err)
		panic(err)

		//return
	}
	Db = dataDase
	Db.AutoMigrate(&Models.User{}, &Models.PetDetail{}, &Models.Dynamics{})
	fmt.Println("链接成功", err)

}
