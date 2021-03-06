package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//文章控制器

func ArticleController(c *gin.Context) {
	if c.Request.Method == "POST" {
		ArticlePost(c)
	} else if c.Request.Method == "GET" {
		ArticleAll(c)
	}
}

func ArticleAll(c *gin.Context) {
	fmt.Println("进入")
	var Article []Models.Article
	if err := Untils.Db.Model(&Models.Article{}).Find(&Article).Error; err != nil {
		Untils.ResponseBadState(c, err)
		return
	}
	Untils.ResponseOkState(c, Article)
}

func ArticlePost(c *gin.Context) {
	FormData := Models.Article{}
	errs := c.Bind(&FormData)
	fmt.Println("绑定有误:", errs)
	if errs != nil {
		Untils.ResponseBadState(c, errs)
	}
	err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		if e1 := tx.Where("user_id=?", FormData.ArticleAuthor).Find(&Models.User{}).Error; e1 != nil {
			return e1
		}
		if e2 := Untils.Db.Model(&Models.Article{}).Create(&FormData).Error; e2 != nil {
			return e2
		}
		return nil
	})
	if err != nil {
		Untils.ResponseBadState(c, err)
		return
	}

	Untils.ResponseOkState(c, FormData)
}
