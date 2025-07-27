package dao

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(128);unique;not null"`
	Password string `gorm:"type:varchar(300);not null"`
	Email    string `gorm:"type:varchar(128);unique;not null"`
}

func (User) TableName() string {
	return "t_user"
}

// 新增用户
func UserInsert(db *gorm.DB, user *User) error {

	err := db.Debug().Create(user).Error
	if err != nil {
		log.Printf("[UserInsert]新增出错！error=%v", err)
	}
	return err

}

// 新增用户
func UserGet(db *gorm.DB, username string) User {

	user := User{}
	db.Debug().Where("username = ?", username).First(&user)
	log.Printf("[UserGet]user=%v", user)
	return user

}
