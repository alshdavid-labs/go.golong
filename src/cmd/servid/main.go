package main

import (
	"golongtemplate/cmd/servid/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	routes.Init()
	routes.Serve()
}
