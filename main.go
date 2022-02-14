package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct {
	Username string `json:"username" binding:"required,min=8,max=20"`
	Password string `json:"password" binding:"required,min=10,max=20"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var login Login
		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "login succeed", "data": login})
	})
	err := router.Run(":8000")
	if err != nil {
		log.Fatal("启动失败，", err)
	}
}
