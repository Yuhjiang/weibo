package routers

import (
	"github.com/Yuhjiang/weibo/controllers"
	"github.com/Yuhjiang/weibo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	userRouter(router)
	articleRouter(router)

	return router
}

func userRouter(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/user/login", controllers.LoginUser)
	// 需要用户登录权限的接口
	auth := router.Group("/user", middleware.JWTAuth())
	auth.GET("", controllers.UserList)
}

func articleRouter(router *gin.Engine) {
	router.GET("/article/:id", controllers.GetArticleById)
	router.GET("/article", controllers.GetArticleList)

	auth := router.Group("/article", middleware.JWTAuth())
	auth.POST("", controllers.CreateArticle)
	auth.PUT("/:id", controllers.UpdateArticle)
	auth.DELETE("/:id", controllers.DeleteArticle)
}
