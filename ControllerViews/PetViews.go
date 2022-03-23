package Views

import (
	"PetService/Conf"
	"PetService/Models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"time"
)

type PetController struct {
}

func (PetController) PetGet(c *gin.Context) {
	//查询宠物
	var PetList []Models.PetDetail
	Conf.Db.Model(&Models.PetDetail{}).Find(&PetList)
	c.JSON(200, gin.H{
		"1":   "sss",
		"res": PetList,
	})
}

func (PetController) PetPost(c *gin.Context) {
	//新增宠物
	petId := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000)
	Pet := Models.PetDetail{}
	err := c.Bind(&Pet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"res": "添加失败",
			"err": err.Error(),
		})
		return
	}
	res := Conf.Db.Transaction(func(tx *gorm.DB) error {
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
		c.JSON(http.StatusBadRequest, gin.H{
			"res": "添加失败",
			"err": res.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "添加成功",
			"res": Pet,
		})
	}

}
