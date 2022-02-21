package controllers

import (
	"github.com/Yuhjiang/weibo/models"
	"github.com/Yuhjiang/weibo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateArticle(c *gin.Context) {
	vUser, _ := c.Get("user")
	user := vUser.(models.User)
	article := models.Article{AuthorId: user.Id}
	err := c.ShouldBind(&article)
	if err != nil {
		utils.ValidateErrorResp(c, err)
		return
	}
	err = models.InsertArticle(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "添加文章失败"})
	}
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func GetArticleById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "not found"})
		return
	}
	article, err := models.GetArticleDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func GetArticleList(c *gin.Context) {
	// 分页查询接口，如果没有传递page和size
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	articles := models.PageArticleList(int(page), int(size))
	c.JSON(http.StatusOK, gin.H{"data": articles})
}
