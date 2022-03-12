package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Dynamics struct {
	gorm.Model
	AuthorID      sql.NullInt64  `gorm:"not null" json:"author_id" binding:"required"`
	DynamicText   sql.NullString `gorm:"not null" json:"dynamic_text" binding:"required"`
	Goods         int64
	DynamicUid    sql.NullInt64 `gorm:"not null;unique;unique_index"`
	Reviews       string        `gorm:"default:'暂无评论'"`
	DynamicImages string        `json:"dynamic_images"`
	DynamicIp     string        `gorm:"not null"`
}

func (Dynamics) TableName() string {
	return "Dynamic"
}
