package user

import "github.com/gin-gonic/gin"

func RouterGroup(Api *gin.RouterGroup) {
	userGroup := Api.Group("/user")
	{
		// 用户注册
		userGroup.POST("/register", func(c *gin.Context) {
			userName := c.DefaultQuery("userName", "guest")
			c.JSON(200, gin.H{
				"message": userName,
			})
		})
	}
}
