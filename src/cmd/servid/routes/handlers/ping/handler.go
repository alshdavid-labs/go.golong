package handlersping

import "github.com/gin-gonic/gin"

// Any is the handler for [Any][/ping]
func Any(c *gin.Context) {
	c.JSON(200, SuccessResponse)
}
