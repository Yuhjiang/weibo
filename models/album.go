package models

import (
	orm "github.com/Yuhjiang/weibo/database"
	"time"
)

type Album struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name" gorm:"required,min=5,max=50"`
	Path       string    `json:"path"`
	Status     int8      `json:"status"`
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"`
}

func (Album) TableName() string {
	return "album"
}

func CreateAlbum(album *Album) error {
	err := orm.DB.Create(album).Error
	return err
}

func GetAlbumList() []Album {
	var albums []Album
	orm.DB.Find(&albums)
	return albums
}
