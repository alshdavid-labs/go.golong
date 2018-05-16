package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

// GetRouter exports the router
func GetRouter() *gin.Engine {
	return r
}

// Serve starts the server
func Serve() {
	address := fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT"))
	fmt.Printf(address)
	r.Run(address)
}
