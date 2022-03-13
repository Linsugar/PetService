package Models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	ArticleAuthor  int64  `gorm:"not null" json:"author" form:"author" binding:"required"`
	ArticleTitle   string `gorm:"not null" json:"title" form:"title" binding:"required"`
	ArticleContent string `gorm:"not null" json:"content" form:"content" binding:"required"`
	ArticleViews   int64  `gorm:"default:0" json:"views" form:"views"`
	ArticleGoods   int64  `gorm:"default:0" json:"goods" form:"goods"`
	ArticleImage   string `gorm:"not null" json:"article_image" form:"article_image" binding:"required"`
}

func (Article) TableName() string {
	return "Article"
}
