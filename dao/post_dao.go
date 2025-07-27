package dao

import (
	"github.com/zhuruiIcarbonx/metaBlog/logger"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string `gorm:"type:varchar(500);not null"`
	Content  string `gorm:"type:longtext;not null"`
	UserId   uint   `gorm:"column:user_id;type:bigint;not null"`
	User     User
	Comments []Comment
}

func (Post) TableName() string {

	return "t_post"
}

func PostCreate(db *gorm.DB, post *Post) error {

	return db.Create(post).Error

}

func PostList(db *gorm.DB, post *Post) []Post {

	list := []Post{}
	if post.Title != "" {
		db.Where("title like '%?%'", post.Title)

	}
	if post.UserId != 0 {
		db.Where("user_id = ?", post.UserId)
	}
	db.Debug().Find(&list)
	logger.Log.Info("list:%v", list)
	return list

}

func PostOne(db *gorm.DB, id int) Post {

	post := Post{}
	db.Debug().Where("id = ?", id).Preload("User").Preload("Comments").First(&post)
	logger.Log.Info("post:%v", post)
	return post

}

func PostUpate(db *gorm.DB, post *Post) error {

	err := db.Debug().Model(Post{}).Where("id = ?", post.ID).Updates(Post{Title: post.Title, Content: post.Content}).Error
	return err
}

func PostDelete(db *gorm.DB, postId int) error {
	return db.Debug().Where("id = ?", postId).Delete(&Post{}).Error
}
