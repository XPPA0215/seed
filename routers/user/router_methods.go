package user

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	userName := c.DefaultQuery("userName", "guest")
	c.JSON(200, gin.H{
		"message": userName,
	})
}
