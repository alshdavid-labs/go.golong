package routes

import (
	"app/services/posts"

	"github.com/kataras/iris"
)

func PostsGet(ctx iris.Context) {
	ctx.JSON(posts_service.Get())
}
