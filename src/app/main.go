package main

import (
	"fmt"
	"os"

	"app/routes"

	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	app.Handle("GET", "/", routes.IndexGet)
	app.Handle("GET", "/posts", routes.PostsGet)
	app.Handle("POST", "/posts", routes.PostsPost)

	app.Run(iris.Addr(setHost()))
}

func setHost() string {
	var PORT = os.Getenv("PORT")
	var ADDRESS = os.Getenv("ADDRESS")

	if PORT == "" {
		PORT = "8080"
	}

	if ADDRESS == "" {
		ADDRESS = "0.0.0.0"
	}

	return fmt.Sprintf("%s:%s", ADDRESS, PORT)
}
