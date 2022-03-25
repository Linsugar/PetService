package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

func PetController(c *gin.Context) {
	if c.Request.Method == "POST" {
		PetPost(c)
	} else if c.Request.Method == "GET" {
		PetGet(c)
	}
}

func PetGet(c *gin.Context) {
	//查询宠物
	var PetList []Models.PetDetail
	err := Untils.Db.Model(&Models.PetDetail{}).Find(&PetList).Error
	if err != nil {
		Untils.ResponseBadState(c, err)
	} else {
		Untils.ResponseOkState(c, PetList)
	}
}

func PetPost(c *gin.Context) {
	//新增宠物
	petId := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000)
	Pet := Models.PetDetail{}
	err := c.Bind(&Pet)
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	res := Untils.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Models.User{}).Where("user_id=?", Pet.PetMaster).Find(&Models.User{}).Error; err != nil {
			return err
		}
		Pet.PetID = petId

		if err2 := tx.Model(&Models.PetDetail{}).Create(&Pet).Error; err2 != nil {
			// 返回任何错误都会回滚事务
			return err2
		}
		// 返回 nil 提交事务
		return nil
	})
	if res != nil {
		Untils.ResponseBadState(c, res)
	} else {
		Untils.ResponseOkState(c, Pet)
	}

}
