package Conf

import (
	"PetService/Models"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Host     string
	Port     string
	Database string
	UserName string
	PassWord string
	CharSet  string
	Addr     string
	Db       *gorm.DB
)

func init() {
	load, err := ini.Load("./Conf/conf.ini")
	if err != nil {
		fmt.Printf("文件加载出错%v\n", err)
		return
	}
	LoadMySql(load)
	LoadRedis(load)
}

func LoadMySql(file *ini.File) {
	//加载Mysql配置文件
	Host = file.Section("MySql").Key("Host").String()
	Port = file.Section("MySql").Key("Port").String()
	Database = file.Section("MySql").Key("Database").String()
	UserName = file.Section("MySql").Key("UserName").String()
	PassWord = file.Section("MySql").Key("PassWord").String()
	CharSet = file.Section("MySql").Key("CharSet").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", UserName, PassWord, Host, Port, Database, CharSet)
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

func LoadRedis(file *ini.File) {
	//加载redis配置文件
	key, err := file.Section("Redis").GetKey("Addr")
	if err != nil {
		return
	}
	Addr = key.String()
}
