package routers

import (
	"github.com/gin-gonic/gin"
	//"leopo/global"
	"leopo/routers/user"
)

func InitWebRouter() *gin.Engine {
	var router = gin.Default()
	Api := router.Group("/api")
	{
		userGroup := Api.Group("/user")
		{
			userGroup.GET("/a", user.A)
		}
	}
	return router
}
