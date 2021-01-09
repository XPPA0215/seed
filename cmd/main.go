package main

import (
	"leopo/routers"
)

func main() {
	router := routers.InitWebRouter()
	//routers.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	_ = router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
