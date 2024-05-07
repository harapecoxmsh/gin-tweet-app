package main

import (
	"gin-tweet-app/controllers"
	"gin-tweet-app/models"
	"gin-tweet-app/repositories"
	"gin-tweet-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	contents := []models.Post{
		{ID: 1, Content: "ginでX風アプリ開発中です"},
		{ID: 2, Content: "頑張ります"},
		{ID: 3, Content: "がんばれえええ"},
	}
	postRepository := repositories.NewPostMemoryRepository(contents)
	postService := services.NewPostService(postRepository)
	postController := controllers.NewPostController(postService)
	r := gin.Default()
	r.GET("/posts", postController.FindAll)
	r.Run("localhost:8080")
}
