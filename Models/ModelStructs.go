package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	//用户信息表
	gorm.Model
	Username string `gorm:"not null"`
	Password string
	Phone string `gorm:"unique;index:phone"`
	CreateCity string
	CreateAddress string
	InitIp string
	NowIp string
	Token string `gorm:"column:token;unique"`
	IsDel bool  //是否删除
	UserId sql.NullString  `gorm:"unique;unique_index;not null"`//不重复id
	InvitePerson int //邀请人id某人为空
	ProfilePicture string //头像地址
}

//自定义表名-默认是结构体名称+s
func (User)TableName()string{
	return "User"
}

type PetDetail struct {
	//宠物资料详细表
	ID int `gorm:"autoIncrement;primaryKey"`
	PetMaster int //宠物主人的id
	PetName string
	PetCall string //联系方式
	Petdetail string //宠物详细介绍
	PetClass string //宠物类型
	PetBuyer int //买主id 默认为空
	PetPhoto string //宠物相册
	PetAvatotr string //宠物头像
	PetVideo string //视频地址
	PetMoney float64 //最初定价
	PetBuyMoney float64 //最终售卖价
}


