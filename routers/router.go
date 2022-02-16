package routers

import (
	"github.com/Yuhjiang/weibo/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	userRouter(router)

	return router
}

func userRouter(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.GET("/user", controllers.UserList)
	router.POST("/user/login", controllers.LoginUser)
}
