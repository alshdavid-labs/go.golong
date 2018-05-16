package handlersindex

import "github.com/gin-gonic/gin"

// Any is the handler for [Any][/]
func Any(c *gin.Context) {
	c.JSON(200, SuccessResponse)
}
