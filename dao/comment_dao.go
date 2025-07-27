package dao

import (
	"github.com/zhuruiIcarbonx/metaBlog/logger"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"type:varchar(500);not null"`
	UserId  uint   `gorm:""column:user_id;type:bigint;not null"`
	PostId  uint   `gorm:""column:post_id;type:bigint;not null"`
}

func (Comment) TableName() string {

	return "t_comment"
}

func CommentCreate(db *gorm.DB, comment *Comment) error {

	return db.Debug().Create(comment).Error

}

func CommentList(db *gorm.DB, postId int) []Comment {

	list := []Comment{}
	db.Debug().Where("post_id = ?", postId).Find(&list).Order(" order by id")
	logger.Log.Printf("list:%v", list)
	return list

}

func CommentDelete(db *gorm.DB, postId int) error {

	err := db.Debug().Where("post_id = ?", postId).Delete(&Comment{}).Error
	return err

}
