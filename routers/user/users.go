package user

import "github.com/gin-gonic/gin"

func A(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
