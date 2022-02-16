package routers

import (
	"github.com/Yuhjiang/weibo/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/register", controllers.RegisterUser)

	return router
}
