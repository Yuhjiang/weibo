package main

import (
	"github.com/Yuhjiang/weibo/routers"
	"log"
)

func main() {
	router := routers.InitRouter()
	err := router.Run(":8000")
	if err != nil {
		log.Fatal("启动失败，", err)
	}
}
