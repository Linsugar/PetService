package Views

import (
	"PetService/Models"
	"PetService/MysqlDo"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

//type Article interface {
//	ArticleAll(c *gin.Context)
//}

//文章控制器
type ArticleController struct {
}

func (ArticleController) ArticleAll(c *gin.Context) {
	fmt.Println("进入")
	var Article []Models.Article
	if err := MysqlDo.Db.Model(&Models.Article{}).Find(&Article).Error; err != nil {
		c.JSON(400, gin.H{
			"msg": "请求错误",
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": &Article,
		"msg":    "请求成功",
	})

}

func (ArticleController) ArticlePost(c *gin.Context) {
	FormData := Models.Article{}
	errs := c.Bind(&FormData)
	fmt.Println("绑定有误:", errs)
	if errs != nil {
		fmt.Println("绑定有误")
		panic(errs)
	}
	err := MysqlDo.Db.Transaction(func(tx *gorm.DB) error {
		if e1 := tx.Where("user_id=?", FormData.ArticleAuthor).Find(&Models.User{}).Error; e1 != nil {
			return e1
		}
		if e2 := MysqlDo.Db.Model(&Models.Article{}).Create(&FormData).Error; e2 != nil {
			return e2
		}
		return nil
	})
	if err != nil {
		c.JSON(400, gin.H{
			"msg":    "参数有误",
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":    "发布成功",
		"result": FormData,
	})
}
