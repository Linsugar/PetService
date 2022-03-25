package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Dynamic struct {
	//必须大写
	AuthorID      int64  `gorm:"not null" json:"author_id" binding:"required" form:"author_id"`
	DynamicText   string `gorm:"not null" json:"dynamic_text" binding:"required" form:"dynamic_text"`
	DynamicImages string `json:"dynamic_images" form:"dynamic_images"`
	DynamicIp     string `json:"dynamicIp" form:"dynamic_ip"`
}

func DynamicAll(c *gin.Context) {
	//获取所有的动态信息
	var allDyadic []Models.Dynamics
	err := Untils.Db.Model(&Models.Dynamics{}).Find(&allDyadic).Error
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"msg":    "请稍后再试",
			"result": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "获取成功",
			"result": allDyadic,
		})
	}

}

func DynamicPost(c *gin.Context) {
	//发布动态
	var bindData Models.Dynamics
	var Data Dynamic
	randID := strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000), 10)
	Rid, _ := strconv.Atoi(randID)

	err := c.Bind(&Data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": err,
		})
		return
	}
	bindData.DynamicIp = c.ClientIP()
	bindData.DynamicUid = sql.NullInt64{Int64: int64(Rid), Valid: true}
	bindData.AuthorID = sql.NullInt64{Int64: Data.AuthorID, Valid: true}
	bindData.DynamicText = sql.NullString{String: Data.DynamicText, Valid: true}
	bindData.DynamicImages = Data.DynamicImages
	ers := Untils.Db.Transaction(func(tx *gorm.DB) error {
		if e1 := tx.Where("user_id=?", &Data.AuthorID).First(&Models.User{}).Error; e1 != nil {
			return e1
		}
		if e2 := tx.Model(&Models.Dynamics{}).Create(&bindData).Error; e2 != nil {
			return e2
		}
		return nil
	})
	if ers != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": ers.Error(),
			"msg": "参数有误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res": bindData,
		"msg": "提交成功",
	})

}
