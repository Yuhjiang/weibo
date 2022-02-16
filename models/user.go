package models

import (
	orm "github.com/Yuhjiang/weibo/database"
	"time"
)

type User struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username" binding:"required,min=5,max=10"`
	Password   string    `json:"password" binding:"required,min=5,max=20"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "user"
}

func InsertUser(user *User) error {
	result := orm.DB.Create(user)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func GetUserByUsername(username string) (User, error) {
	user := User{}
	res := orm.DB.Where("name = ?", username).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
