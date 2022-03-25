package Conf

import (
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

var (
	Host         string
	Port         string
	Database     string
	UserName     string
	PassWord     string
	CharSet      string
	Addr         string
	Secret       string
	Appid        string
	OpenConns    int
	MaxIdleConns int
	AccessKey    string
	SecretKey    string
)

func init() {
	load, err := ini.Load("./Conf/conf.ini")
	if err != nil {
		fmt.Printf("文件加载出错%v\n", err)
		return
	}
	LoadMySql(load)
	LoadRedis(load)
	LoadWeiXinKey(load)
}

func LoadMySql(file *ini.File) {
	//加载Mysql配置文件
	Host = file.Section("MySql").Key("Host").String()
	Port = file.Section("MySql").Key("Port").String()
	Database = file.Section("MySql").Key("Database").String()
	UserName = file.Section("MySql").Key("UserName").String()
	PassWord = file.Section("MySql").Key("PassWord").String()
	CharSet = file.Section("MySql").Key("CharSet").String()
	OpenConns, _ = strconv.Atoi(file.Section("Mysql").Key("OpenConns").String())
	MaxIdleConns, _ = strconv.Atoi(file.Section("Mysql").Key("MaxIdleConns").String())
}

func LoadRedis(file *ini.File) {
	//加载redis配置文件
	key, err := file.Section("Redis").GetKey("Addr")
	if err != nil {
		return
	}
	Addr = key.String()
}

func LoadWeiXinKey(file *ini.File) {
	Appid = file.Section("WeiXinKey").Key("Appid").String()
	Secret = file.Section("WeiXinKey").Key("Secret").String()
}

func LoadQiNiuKey(file *ini.File) {
	AccessKey = file.Section("QiNiu").Key("AccessKey").String()
	SecretKey = file.Section("QiNiu").Key("SecretKey").String()
}
