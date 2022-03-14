package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	//用户信息表
	gorm.Model
	Username       string `gorm:"not null" json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Phone          string `gorm:"unique;index:phone" json:"phone" binding:"required"`
	CreateCity     string `gorm:"default:'成都'" json:"create_city" form:"create_city"`
	CreateAddress  string `gorm:"default:'成都高新区'"`
	InitIp         string
	NowIp          string
	Token          string         `gorm:"column:token;"`
	IsDel          bool           `gorm:"column:isdel;default:false"`                                                          //是否删除
	UserId         sql.NullString `gorm:"unique;unique_index;not null"`                                                        //不重复id
	InvitePerson   int            `gorm:"default:'6666'" json:"invitePerson" form:"invitePerson"`                              //邀请人id某人为空
	ProfilePicture string         `gorm:"default:'http://cdn.tlapp.club/pet.png'" json:"profilePicture" form:"profilePicture"` //头像地址
}

//自定义表名-默认是结构体名称+s
func (User) TableName() string {
	return "User"
}

type PetDetail struct {
	//宠物资料详细表
	gorm.Model
	PetID       int64         `grom:"not null;unique;index:pet"`
	PetMaster   string        `gorm:"not null" json:"petMaster" binding:"required" form:"petMaster"` //宠物主人的id
	PetName     string        `json:"pet_name" form:"pet_name" binding:"required" `
	PetCall     string        `gorm:"default:'无'" json:"petCall" form:"petCall"`                                                        //联系方式
	Petdetail   string        `gorm:"default:'暂无介绍'" json:"petdetail" form:"petdetail"`                                                 //宠物详细介绍
	PetClass    string        `gorm:"not null;default:'0'" json:"petClass" form:"petClass"`                                             //宠物类型
	PetBuyer    sql.NullInt32 `json:"petBuyer" form:"PetBuyer"`                                                                         //买主id 默认为空
	PetPhoto    string        `gorm:"default:'暂无'" json:"petPhoto" form:"petPhoto"`                                                     //宠物相册
	PetAvatotr  string        `gorm:"default:'http://cdn.tlapp.club/pet.png'" form:"pet_avatotr" json:"pet_avatotr" binding:"required"` //宠物头像
	PetVideo    string        `gorm:"default:'暂无'" json:"pet_video" form:"pet_video"`                                                   //视频地址
	PetMoney    float64       `gorm:"default:'0.0'" json:"petMoney" form:"petMoney"`                                                    //最初定价
	PetBuyMoney float64       `gorm:"default:'0.0'" json:"petBuyMoney" form:"petBuyMoney"`                                              //最终售卖价
	PetContent  string        `json:"petContent" form:"petContent" binding:"required"`                                                  //最终售卖价
	PetAge      float64       `json:"petAge" form:"petAge" binding:"required"`                                                          //最终售卖价
	PetGender   string        `gorm:"defalut:'MALE'" json:"petGender" form:"petGender"`                                                 //最终售卖价
	PetWeight   float32       `gorm:"default:'0.0'" json:"petWeight" form:"petWeight"`                                                  //最终售卖价
	PetLocation string        `gorm:"defalut:'[10.0,20]'" json:"petLocation" form:"petLocation"`
}
