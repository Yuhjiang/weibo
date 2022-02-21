package models

import (
	orm "github.com/Yuhjiang/weibo/database"
	"time"
)

type Article struct {
	Id         int64         `json:"id" gorm:"primaryKey"`
	AuthorId   int64         `json:"authorId" binding:"required"`
	Title      string        `json:"title" binding:"required,min=5,max=30"`
	Tags       string        `json:"tags" binding:"required,min=1,max=30"`
	Short      string        `json:"short" binding:"required,max=255"`
	CreateTime time.Time     `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time     `json:"updateTime" gorm:"autoUpdateTime"`
	Detail     ArticleDetail `json:"detail" gorm:"foreignKey:ArticleId"`
}

func (Article) TableName() string {
	return "article"
}

type ArticleDetail struct {
	Id        int64  `json:"id"`
	ArticleId int64  `json:"articleId"`
	Content   string `json:"content" binding:"required"`
}

func (ArticleDetail) TableName() string {
	return "article_detail"
}

func InsertArticle(article *Article) error {
	err := orm.DB.Create(article).Error
	return err
}

func GetArticleDetail(id int64) (Article, error) {
	article := Article{}
	article.Id = id
	res := orm.DB.Joins("Detail").First(&article).Error
	if res != nil {
		return article, res
	} else {
		return article, nil
	}
}

type ArticleVO struct {
	Id         int64     `json:"id" gorm:"primaryKey"`
	AuthorId   int64     `json:"authorId"`
	AuthorName string    `json:"authorName"`
	Title      string    `json:"title"`
	Tags       string    `json:"tags"`
	Short      string    `json:"short"`
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"`
}

func GetArticleList() []ArticleVO {
	var articles []ArticleVO
	orm.DB.Model(&Article{}).Select(
		"article.id, author_id, user.username AS author_name, title, tags, short, " +
			"article.create_time").Joins("LEFT JOIN user on user.id = article.author_id").Find(&articles)
	return articles
}

type PageArticle struct {
	Data  []ArticleVO `json:"data"`
	Count int64       `json:"count"`
}

// PageArticleList 分页查询的文章列表
func PageArticleList(page, pageSize int) PageArticle {
	var articles []ArticleVO
	tx := orm.DB.Begin()
	defer tx.Commit()
	tx.Model(&Article{}).Select(
		"article.id, author_id, user.username AS author_name, title, tags, short, " +
			"article.create_time").Joins(
		"LEFT JOIN user on user.id = article.author_id").Offset(
		(page - 1) * pageSize).Limit(pageSize).Find(&articles)
	var count int64
	tx.Model(&Article{}).Count(&count)
	return PageArticle{Data: articles, Count: count}
}
