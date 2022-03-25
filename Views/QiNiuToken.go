package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type GetToken struct {
	UserId string `json:"userId" form:"userId" binding:"required"`
}

func SetQINiuToken(c *gin.Context) {
	Data := GetToken{}
	err := c.ShouldBind(&Data)
	if err != nil {
		Untils.ResponseBadState(c, err)
	}
	result := Untils.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Models.User{}).Where("user_id=?", Data.UserId).First(&Models.User{}).Error; err != nil {
			return err
		}
		return nil
	})
	if result != nil {
		Untils.ResponseBadState(c, result)
		return
	}
	token := Untils.QiNiuToken()
	Untils.ResponseOkState(c, token)
}
