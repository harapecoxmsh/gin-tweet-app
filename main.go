package main

import (
	"gin-tweet-app/controllers"
	"gin-tweet-app/infra"
	"gin-tweet-app/repositories"
	"gin-tweet-app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository)
	postController := controllers.NewPostController(postService)
	r := gin.Default()
	r.Use(cors.Default())
	itemRouter := r.Group("/posts")
	itemRouter.GET("", postController.FindAll)
	itemRouter.GET("/:id", postController.FindByID)
	itemRouter.POST("", postController.Create)
	itemRouter.DELETE("/:id", postController.Delete)

	r.Run("localhost:8080")
}
