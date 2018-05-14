package routes

import (
	"app/models"
	"app/services/posts"

	"github.com/kataras/iris"
)

func PostsPost(ctx iris.Context) {
	var post models.Post
	var err = ctx.ReadJSON(&post)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"message": "Invalid Request",
			"error":   err.Error(),
		})
		return
	}

	if post.Content == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"message": "Invalid Request",
			"error":   "Missing field",
		})
		return
	}

	var created = posts_service.Add(post.Content)

	ctx.JSON(iris.Map{
		"message": "Post Created",
		"post":    created,
	})
}
