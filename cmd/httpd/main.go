package main

import (
	"fmt"

	"golong/cmd/httpd/handlers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	router.Any("/", handlers.IndexGet)
	router.Any("/ping", handlers.PingAny)

	router.Run("0.0.0.0:5000")
}
