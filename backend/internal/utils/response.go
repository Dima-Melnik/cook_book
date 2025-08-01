package utils

import "github.com/gin-gonic/gin"

func SendResponseError(c *gin.Context, status int, err any) {
	c.JSON(status, gin.H{"error": err})
	c.Abort()
}

func SendResponseJSON(c *gin.Context, status int, resp any) {
	c.JSON(status, resp)
}
