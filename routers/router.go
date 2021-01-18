package routers

import (
	"github.com/gin-gonic/gin"
	"leopo/routers/user"
)

func InitWebRouter() *gin.Engine {
	var router = gin.Default()
	Api := router.Group("/api")
	// 账户类路由
	user.RouterGroup(Api)
	return router
}
