package controllers

import (
	"github.com/Yuhjiang/weibo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func UploadAlbum(c *gin.Context) {
	f, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "图片上传失败"})
		return
	}
	ext := filepath.Ext(f.Filename)
	var fType string
	switch ext {
	case ".jpg", ".png", ".gif", ".jpeg":
		fType = "img"
	default:
		fType = "other"
	}
	if fType == "other" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "图片格式不正确"})
		return
	}
	path := filepath.Join("./static/album", f.Filename)
	err = c.SaveUploadedFile(f, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "图片保存失败"})
		return
	}
	album := models.Album{
		Name: f.Filename,
		Path: path,
	}
	err = models.CreateAlbum(&album)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "图片保存失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": album})
	}
}

func GetAlbumList(c *gin.Context) {
	albums := models.GetAlbumList()
	c.JSON(http.StatusOK, gin.H{"data": albums})
}
