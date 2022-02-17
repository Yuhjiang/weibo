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
	_, exist := models.GetUserByUsername(user.Username)
	if exist {
		c.JSON(http.StatusOK, gin.H{"msg": "用户名已存在"})
		return
	}
	err = models.InsertUser(&user)
	if err != nil {
		log.Fatal("注册用户失败", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UserList(c *gin.Context) {
	users := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func LoginUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		utils.ValidateErrorResp(c, err)
		return
	}
	login := models.LoginUser(&user)
	if login {
		token, _ := models.CreateToken(user)
		c.JSON(http.StatusOK, gin.H{"data": user, "msg": "登录成功",
			"token": token})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "账号或密码错误"})
	}
}
