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
