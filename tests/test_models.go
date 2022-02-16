package main

import (
	"fmt"
	orm "github.com/Yuhjiang/weibo/database"
	"github.com/Yuhjiang/weibo/models"
)

func main() {
	user := models.User{}
	res := orm.DB.Where("username = ?", "yuhao1").First(&user)
	fmt.Println(user, res.Error)
}
