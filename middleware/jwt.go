package middleware

import (
	"github.com/Yuhjiang/weibo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const TokenUnValid = "请求令牌无效"

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": TokenUnValid})
			c.Abort()
			return
		}
		user, err := models.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": TokenUnValid})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func GetCurrentUser(c *gin.Context) (models.User, bool) {
	var user models.User
	vUser, exist := c.Get("user")
	if !exist {
		return user, exist
	}
	user, ok := vUser.(models.User)
	if !ok {
		return user, ok
	} else {
		return user, true
	}
}
