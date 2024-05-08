package main

import (
	"gin-tweet-app/controllers"
	"gin-tweet-app/infra"
	"gin-tweet-app/repositories"
	"gin-tweet-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository)
	postController := controllers.NewPostController(postService)
	r := gin.Default()
	itemRouter := r.Group("/posts")
	itemRouter.GET("", postController.FindAll)
	itemRouter.POST("", postController.Create)

	r.Run("localhost:8080")
}
