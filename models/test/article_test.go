package test

import (
	"github.com/Yuhjiang/weibo/models"
	"testing"
)

func TestGetArticleDetail(t *testing.T) {
	article, err := models.GetArticleDetail(2)
	if err != nil {
		t.Error(err)
	}
	t.Log(article)
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		AuthorId: 2,
		Title:    "forth article",
		Tags:     "test",
		Short:    "short description",
		Detail: models.ArticleDetail{
			Content: "hello",
		},
	}
	err := models.InsertArticle(&article)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(article)
	}
}

func TestGetArticleList(t *testing.T) {
	articles := models.GetArticleList()
	for _, a := range articles {
		t.Log(a)
	}
}
