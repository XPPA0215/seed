package routers

import (
	"github.com/gin-gonic/gin"
	//"leopo/global"
	"leopo/routers/users"
)

func InitWebRouter() *gin.Engine {
	var router = gin.Default()
	Api := router.Group("/api")
	Api.GET("/users", users.A)
	return router
}
