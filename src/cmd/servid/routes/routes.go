package routes

import (
	"golongtemplate/cmd/servid/routes/handlers/index"
	"golongtemplate/cmd/servid/routes/handlers/ping"
)

// Init the routes
func Init() {
	r.Any("", handlersindex.Any)
	r.Any("/ping", handlersping.Any)
}
