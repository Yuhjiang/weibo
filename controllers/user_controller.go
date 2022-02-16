package controllers

import (
	"github.com/Yuhjiang/weibo/models"
	"github.com/Yuhjiang/weibo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		utils.ValidateErrorResp(c, err)
		return
	}
	_, notFound := models.GetUserByUsername(user.Username)
	if notFound == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "用户名已存在"})
		return
	}
	err = models.InsertUser(&user)
	if err != nil {
		log.Fatal("注册用户失败", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
