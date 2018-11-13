package handlers

import (
	"github.com/gin-gonic/gin"
)

func IndexGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
}
