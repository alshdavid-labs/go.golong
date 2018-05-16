package handlersping

import "github.com/gin-gonic/gin"

// SuccessResponse is the response for [Any][/ping]
var SuccessResponse = gin.H{
	"message": "pong",
}
